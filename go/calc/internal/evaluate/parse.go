package evaluate

import (
	"strconv"
)

type parser struct {
	input     chan item
	output    chan node
	token     [1]item
	peekCount int
}

func parse(input chan item) *parser {
	p := &parser{
		input:  input,
		output: make(chan node),
	}
	go p.parse()
	return p
}

func (p *parser) peek() item {
	if p.peekCount > 0 {
		return p.token[p.peekCount-1]
	}
	p.peekCount = 1
	p.token[0] = <-p.input
	return p.token[0]
}

func (p *parser) next() item {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		p.token[0] = <-p.input
	}
	return p.token[p.peekCount]
}

func (p *parser) parse() {
	for p.peek().typ != itemEOF {
		exp := p.expression()
		p.output <- exp

		// Skip to next eol for recovering error.
		if exp.Type() == nodeError {
			for p.peek().typ != itemEOL && p.peek().typ != itemEOF {
				// TODO: output as error
				p.next()
			}
		}

		// Skip EOL
		if p.peek().typ == itemEOL {
			p.next()
		}
	}
	close(p.output)
}

func (p *parser) expression() node {
	t1 := p.term()
	if t1.Type() == nodeError {
		return t1
	}
	for {
		if p.peek().typ == itemAdd || p.peek().typ == itemSub {
			op := p.next()
			var t binaryOpType
			switch op.typ {
			case itemAdd:
				t = binaryOpAdd
			case itemSub:
				t = binaryOpSub
			default:
				panic("unknown itemType")
			}
			t2 := p.term()
			if t2.Type() == nodeError {
				return t2
			}
			// FIXME: Fix to left-to-right associative.
			t1 = newBinaryOpNode(t1, t2, t)
		} else {
			break
		}
	}
	return t1
}

func (p *parser) term() node {
	e1 := p.primaryExpression()
	if e1.Type() == nodeError {
		return e1
	}
	for {
		if p.peek().typ == itemMul || p.peek().typ == itemDiv {
			op := p.next()
			var t binaryOpType
			switch op.typ {
			case itemMul:
				t = binaryOpMul
			case itemDiv:
				t = binaryOpDiv
			default:
				panic("unknown itemType")
			}
			e2 := p.primaryExpression()
			if e2.Type() == nodeError {
				return e2
			}
			// FIXME: Fix to left-to-right associative.
			e1 = newBinaryOpNode(e1, e2, t)
		} else {
			break
		}
	}
	return e1
}

func (p *parser) primaryExpression() node {
	isMinus := false

	if p.peek().typ == itemVariable {
		v := p.next()
		if p.peek().typ != itemAssign {
			return newVariableRefNode(v.val)
		}
		p.next()
		e := p.expression()
		if e.Type() == nodeError {
			return e
		}
		return newAssignNode(v.val, e)
	}
	if p.peek().typ == itemSub {
		isMinus = true
		p.next()
	}

	n := p.next()

	// TODO support functions with more than one parameter
	if n.typ == itemFunction {
		p1 := p.next()

		if p1.typ == itemLParen {

			n2 := p.expression()
			if n2.Type() == nodeError {
				return n2
			}

			n3 := p.next()
			if n3.typ != itemRParen {
				return newErrorNode("expect RParen")
			}
			return newFunctionNode(n2, n.val)

		}
		return newErrorNode("expect LParen")

	} else if n.typ == itemLParen {
		n2 := p.expression()
		if n2.Type() == nodeError {
			return n2
		}
		n3 := p.next()
		if n3.typ != itemRParen {
			return newErrorNode("expect RParen")
		}
		if !isMinus {
			return n2
		}
		return newUnaryOpNode(n2, unaryOpMinus)
	} else if n.typ != itemDoubleLiteral {
		return newErrorNode("unexpected item")
	}
	v, err := strconv.ParseFloat(n.val, 64)
	if err != nil {
		return newErrorNode("unexpected value " + err.Error())
	}

	if !isMinus {
		return newValueNode(v)
	}
	return newUnaryOpNode(newValueNode(v), unaryOpMinus)
}
