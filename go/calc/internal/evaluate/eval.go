package evaluate

import (
	"errors"
	"fmt"
	"math"
)

type function struct {
	args      int
	operation func(args []float64) float64
}

var (
	errUnknownFunc   = errors.New("unknown function")
	errUnassignedVar = errors.New("unassigned variable")
	errUnknownType   = errors.New("unknown type")
	errUnknownBin    = errors.New("unknown binaryOpType")
	errUnknownOpType = errors.New("unknown unaryOpType")
	errNode          = errors.New("node error")
	errWrongType     = errors.New("wrong type")
)

var functions = map[string]function{
	"abs": {
		args: 1,
		operation: func(args []float64) float64 {
			return math.Abs(args[0])
		},
	},
	"floor": {
		args: 1,
		operation: func(args []float64) float64 {
			return math.Floor(args[0])
		},
	},
	"round": {
		args: 1,
		operation: func(args []float64) float64 {
			return math.Round(args[0])
		},
	},
}

type env struct {
	global map[string]value
}

func newEnv() env {
	return env{
		global: make(map[string]value),
	}
}

func (e env) eval(n node) value {
	switch n.Type() {
	case nodeBinaryOp:
		nn, ok := n.(*binaryOpNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		left := e.eval(nn.lhs)

		if left.err != nil {
			return left
		}

		right := e.eval(nn.rhs)
		if right.err != nil {
			return right
		}

		switch nn.opTyp {
		case binaryOpAdd:
			return value{v: left.v + right.v, err: nil}
		case binaryOpSub:
			return value{v: left.v - right.v, err: nil}
		case binaryOpMul:
			return value{v: left.v * right.v, err: nil}
		case binaryOpDiv:
			return value{v: left.v / right.v, err: nil}
		default:
			return value{v: 0, err: errUnknownBin}
		}

	case nodeUnaryOp:
		nn, ok := n.(*unaryOpNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		v := e.eval(nn.operand)

		if v.err != nil {
			return v
		}

		switch nn.opTyp {
		case unaryOpMinus:
			return value{v: -v.v, err: nil}
		default:
			return value{v: 0, err: errUnknownOpType}
		}

	case nodeValue:
		nn, ok := n.(*valueNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		return value{v: nn.v, err: nil}

	case nodeError:
		nn, ok := n.(*errorNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		return value{v: 0, err: fmt.Errorf("%w, %s", errNode, nn.err)}

	case nodeFunction:
		nn, ok := n.(*functionNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		v := e.eval(nn.operand)
		if v.err != nil {
			return v
		}

		if _, exists := functions[nn.f]; !exists {
			return value{v: 0, err: fmt.Errorf("%w, %s", errUnknownFunc, nn.f)}
		}

		args := []float64{v.v}

		return value{v: functions[nn.f].operation(args), err: nil}

	// case nodeAssign:
	// 	nn := n.(*assignNode)
	// 	v := e.eval(nn.expr)
	// 	if v.err != nil {
	// 		return v
	// 	}
	// 	e.global[nn.variable] = v
	// 	return value{v: v.v, err: nil}
	case nodeVariableRef:
		nn, ok := n.(*variableRefNode)
		if !ok {
			return value{v: 0, err: errWrongType}
		}

		v, ok := e.global[nn.variable]
		if !ok {
			return value{v: 0, err: fmt.Errorf("%w, %s", errUnassignedVar, nn.variable)}
		}

		return value{v: v.v, err: nil}
	default:
		return value{v: 0, err: errUnknownType}
	}
}
