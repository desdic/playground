package evaluate

import (
	"fmt"
)

type value struct {
	v   float64
	err error
}

type node interface {
	Type() nodeType
	String() string
}

type nodeType int

const (
	nodeBinaryOp nodeType = iota
	nodeUnaryOp
	nodeValue
	nodeError
	nodeAssign
	nodeFunction
	nodeVariableRef
)

type binaryOpType int

const (
	binaryOpAdd binaryOpType = iota
	binaryOpSub
	binaryOpMul
	binaryOpDiv
)

type binaryOpNode struct {
	opTyp    binaryOpType
	lhs, rhs node
}

func newBinaryOpNode(lhs node, rhs node, op binaryOpType) *binaryOpNode {
	return &binaryOpNode{
		opTyp: op,
		lhs:   lhs,
		rhs:   rhs,
	}
}

func (n *binaryOpNode) Type() nodeType {
	return nodeBinaryOp
}

func (n *binaryOpNode) String() string {
	var opStr string
	switch n.opTyp {
	case binaryOpAdd:
		opStr = "+"
	case binaryOpSub:
		opStr = "-"
	case binaryOpMul:
		opStr = "*"
	case binaryOpDiv:
		opStr = "/"
	default:
		opStr = fmt.Sprintf("%d", n.opTyp)
	}
	return fmt.Sprintf("(%s %s %s)", opStr, n.lhs.String(), n.rhs.String())
}

type unaryOpType int

const (
	unaryOpMinus unaryOpType = iota
)

type unaryOpNode struct {
	opTyp   unaryOpType
	operand node
}

func newUnaryOpNode(operand node, op unaryOpType) *unaryOpNode {
	return &unaryOpNode{
		opTyp:   op,
		operand: operand,
	}
}

func (n *unaryOpNode) Type() nodeType {
	return nodeUnaryOp
}

func (n *unaryOpNode) String() string {
	var opStr string
	switch n.opTyp {
	case unaryOpMinus:
		opStr = "-"
	default:
		opStr = fmt.Sprintf("%d", n.opTyp)
	}
	return fmt.Sprintf("(%s %s)", opStr, n.operand.String())
}

type functionNode struct {
	f       string
	operand node
}

func newFunctionNode(operand node, f string) *functionNode {
	return &functionNode{
		f:       f,
		operand: operand,
	}
}

func (n *functionNode) Type() nodeType {
	return nodeFunction
}

func (n *functionNode) String() string {
	return fmt.Sprintf("(%s)", n.operand.String())
}

type valueNode struct {
	v float64
}

func newValueNode(v float64) *valueNode {
	return &valueNode{
		v: v,
	}
}

func (n *valueNode) Type() nodeType {
	return nodeValue
}

func (n *valueNode) String() string {
	return fmt.Sprintf("%f", n.v)
}

type errorNode struct {
	err string
}

func newErrorNode(err string) *errorNode {
	return &errorNode{
		err: err,
	}
}

func (n *errorNode) Type() nodeType {
	return nodeError
}

func (n *errorNode) String() string {
	return n.err
}

type assignNode struct {
	variable string
	expr     node
}

func newAssignNode(v string, e node) *assignNode {
	return &assignNode{variable: v, expr: e}
}

func (a *assignNode) Type() nodeType {
	return nodeAssign
}

func (a *assignNode) String() string {
	return fmt.Sprint(a.variable, "=")
}

type variableRefNode struct {
	variable string
}

func newVariableRefNode(v string) *variableRefNode {
	return &variableRefNode{variable: v}
}

func (v *variableRefNode) Type() nodeType {
	return nodeVariableRef
}

func (v *variableRefNode) String() string {
	return v.variable
}
