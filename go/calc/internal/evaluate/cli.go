package evaluate

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errNoInput         = errors.New("no input")
	errReservedKeyword = errors.New("is reserved keyword cannot be used as a variable")
)

// Evaluate TODO.
func Evaluate(str string, vars map[string]float64) (f float64, err error) {
	l := lex(strings.NewReader(str))
	p := parse(l.items)
	e := newEnv()

	for k, v := range vars {
		if _, exists := functions[k]; exists {
			return 0, fmt.Errorf("%s %w", k, errReservedKeyword)
		}

		e.global[k] = value{v: v, err: nil}
	}

	for n := range p.output {
		v := e.eval(n)

		return v.v, v.err
	}

	return 0.0, errNoInput
}
