package evaluate

import (
	"bytes"
	"testing"
)

type lexTest struct {
	name  string
	input string
	items []item
}

var origin = pos{line: 1, col: 1}

var tEOF1 = item{itemEOF, origin, ""}
var tEOF2 = item{itemEOF, pos{line: 1, col: 2}, ""}
var tEOF3 = item{itemEOF, pos{line: 1, col: 3}, ""}
var tEOF4 = item{itemEOF, pos{line: 1, col: 4}, ""}

var lexTests = []lexTest{
	{"empty", "", []item{tEOF1}},
	{"space", " ", []item{item{itemEOF, pos{line: 1, col: 2}, ""}}},
	{"number1", "1", []item{item{itemDoubleLiteral, origin, "1"}, tEOF2}},
	{"number11", "11", []item{item{itemDoubleLiteral, origin, "11"}, tEOF3}},
	{"number0", "0", []item{item{itemDoubleLiteral, origin, "0"}, tEOF2}},
	{"number1.0", "1.0", []item{item{itemDoubleLiteral, origin, "1.0"}, tEOF4}},
	{"number1.", "1.", []item{item{itemError, origin, "digit not appear next to dot"}}},
	{"number01", "01", []item{item{itemDoubleLiteral, origin, "01"}, tEOF3}},
	{".", ".", []item{item{itemError, origin, "unexpected character"}}},
	{"i", "i", []item{item{itemVariable, origin, "i"}, tEOF2}},
	{"two number", "1 2", []item{
		item{itemDoubleLiteral, origin, "1"},
		item{itemDoubleLiteral, pos{line: 1, col: 3}, "2"},
		item{itemEOF, pos{line: 1, col: 4}, ""},
	}},
	{"add", "1+2", []item{
		item{itemDoubleLiteral, origin, "1"},
		item{itemAdd, pos{line: 1, col: 2}, "+"},
		item{itemDoubleLiteral, pos{line: 1, col: 3}, "2"},
		tEOF4}},
	{"sub", "1-2", []item{
		item{itemDoubleLiteral, origin, "1"},
		item{itemSub, pos{line: 1, col: 2}, "-"},
		item{itemDoubleLiteral, pos{line: 1, col: 3}, "2"},
		tEOF4}},
	{"mul", "1*2", []item{
		item{itemDoubleLiteral, origin, "1"},
		item{itemMul, pos{line: 1, col: 2}, "*"},
		item{itemDoubleLiteral, pos{line: 1, col: 3}, "2"},
		tEOF4}},
	{"div", "1/2", []item{
		item{itemDoubleLiteral, origin, "1"},
		item{itemDiv, pos{line: 1, col: 2}, "/"},
		item{itemDoubleLiteral, pos{line: 1, col: 3}, "2"},
		tEOF4}},
	{"paren", "(1)", []item{
		item{itemLParen, origin, "("},
		item{itemDoubleLiteral, pos{line: 1, col: 2}, "1"},
		item{itemRParen, pos{line: 1, col: 3}, ")"},
		tEOF4}},
	//{"assign", "a=1", []item{
	//item{itemVariable, origin, "a"},
	//item{itemAssign, pos{line: 1, col: 2}, "="},
	//item{itemDoubleLiteral, pos{line: 1, col: 3}, "1"},
	//tEOF4}},
	// TODO test functions
}

func collect(t *lexTest) (items []item) {
	buf := bytes.NewBufferString(t.input)
	l := lex(buf)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func equal(i1, i2 []item) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].typ != i2[k].typ {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
		if i1[k].pos != i2[k].pos {
			return false
		}
	}
	return true
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		items := collect(&test)
		if !equal(items, test.items) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}

type stringTest struct {
	name  string
	input item
	str   string
}

var stringTests = []stringTest{
	{"add", item{itemAdd, origin, "+"}, `add:"+"(1,1)`},
	{"sub", item{itemSub, origin, "-"}, `sub:"-"(1,1)`},
	{"mul", item{itemMul, origin, "*"}, `mul:"*"(1,1)`},
	{"div", item{itemDiv, origin, "/"}, `div:"/"(1,1)`},
	{"assign", item{itemAssign, origin, "="}, `assign:"="(1,1)`},
	{"doubleLiteral", item{itemDoubleLiteral, origin, "1.0"}, `doubleLiteral:"1.0"(1,1)`},
	{"eol", item{itemEOL, origin, ""}, `eol:""(1,1)`},
	{"eof", item{itemEOF, origin, ""}, `eof:""(1,1)`},
	{"error", item{itemError, origin, "error"}, `error:"error"(1,1)`},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		if test.input.String() != test.str {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, test.input, test.str)
		}
	}
}
