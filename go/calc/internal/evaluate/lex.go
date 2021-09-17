package evaluate

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

type item struct {
	typ itemType
	pos pos
	val string
}

type pos struct {
	line, col int
}

func (i item) String() string {
	return fmt.Sprintf("%s:%q(%d,%d)", i.typ, i.val, i.pos.line, i.pos.col)
}

type itemType int

const (
	itemAdd itemType = iota
	itemSub
	itemMul
	itemDiv
	itemFunction
	itemAssign
	itemDoubleLiteral
	itemVariable
	itemLParen
	itemRParen
	itemEOL
	itemEOF
	itemError
)

const eof = -1

func (t itemType) String() string {
	switch t {
	case itemAdd:
		return "add"
	case itemSub:
		return "sub"
	case itemMul:
		return "mul"
	case itemDiv:
		return "div"
	case itemFunction:
		return "function"
	case itemAssign:
		return "assign"
	case itemDoubleLiteral:
		return "doubleLiteral"
	case itemVariable:
		return "variable"
	case itemLParen:
		return "lparen"
	case itemRParen:
		return "rparen"
	case itemEOF:
		return "eof"
	case itemEOL:
		return "eol"
	case itemError:
		return "error"
	default:
		panic(errors.New("unexpected itemType"))
	}
}

type stateFn func(*lexer) stateFn

type lexer struct {
	input  *bufio.Reader
	buffer bytes.Buffer
	state  stateFn
	pos    pos
	start  pos
	items  chan item
}

func (l *lexer) nextItem() item {
	item := <-l.items

	return item
}

func lex(input io.Reader) *lexer {
	l := &lexer{
		input: bufio.NewReader(input),
		items: make(chan item),
		pos:   pos{line: 1, col: 1},
		start: pos{line: 1, col: 1},
	}

	go l.run()

	return l
}

func (l *lexer) run() {
	for l.state = lexInitial; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) next() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	if r != '\n' {
		l.pos.col += w
	} else {
		l.pos.line++
		l.pos.col = 1
	}

	l.buffer.WriteRune(r)
	return r
}

func (l *lexer) peek() rune {
	lead, err := l.input.Peek(1)
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.errorf("%s", err.Error())
		return 0
	}

	p, err := l.input.Peek(runeLen(lead[0]))
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.errorf("%s", err.Error())
		return 0
	}
	r, _ := utf8.DecodeRune(p)
	return r
}

func runeLen(lead byte) int {
	if lead < 0xC0 {
		return 1
	} else if lead < 0xE0 {
		return 2
	} else if lead < 0xF0 {
		return 3
	}
	return 4
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.buffer.String()}
	l.start = l.pos
	l.buffer.Truncate(0)
}

func (l *lexer) truncate() {
	l.start = l.pos
	l.buffer.Truncate(0)
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.peek()) >= 0 {
		l.next()
		return true
	}
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.peek()) >= 0 {
		l.next()
	}
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func (l *lexer) hasPrefix(prefix string) bool {
	p, err := l.input.Peek(len(prefix))
	if err == io.EOF {
		return false
	} else if err != nil {
		l.errorf("%v", err.Error())
		return false
	}
	return string(p) == prefix
}

//func (l *lexer) nextRuneCount(count int) {
//for i := 0; i < count; i++ {
//l.next()
//}
//}

var alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func lexInitial(l *lexer) stateFn {
LOOP:
	for {
		r := l.peek()
		switch r {
		case ' ':
			l.next()
			l.truncate()
		case '+':
			l.next()
			l.emit(itemAdd)
		case '-':
			l.next()
			l.emit(itemSub)
		case '*':
			l.next()
			l.emit(itemMul)
		case '/':
			l.next()
			l.emit(itemDiv)
		// case '=':
		// l.next()
		// l.emit(itemAssign)
		case '(':
			l.next()
			l.emit(itemLParen)
		case ')':
			l.next()
			l.emit(itemRParen)
		// case '\n':
		// l.next()
		// l.emit(itemEOL)
		case eof:
			l.next()
			break LOOP
		default:
			if strings.IndexRune("0123456789", r) >= 0 {
				return lexDoubleLiteral
			}
			if strings.IndexRune(alphabets, r) >= 0 {
				return lexVariable
			}
			l.errorf("unexpected character")
			l.next()
		}
	}
	l.emit(itemEOF)
	return nil
}

func lexDoubleLiteral(l *lexer) stateFn {
	if !l.accept("0123456789") {
		return l.errorf("bad digit for number")
	}
	l.acceptRun("0123456789")
	if l.accept(".") {
		if !l.accept("0123456789") {
			return l.errorf("digit not appear next to dot")
		}
		l.acceptRun("0123456789")
	}
	l.emit(itemDoubleLiteral)
	return lexInitial
}

func lexVariable(l *lexer) stateFn {
	if !l.accept(alphabets) {
		return l.errorf("bad variable")
	}
	l.acceptRun(alphabets)

	if _, exists := functions[l.buffer.String()]; exists {
		l.emit(itemFunction)
	} else {
		l.emit(itemVariable)
	}

	return lexInitial
}
