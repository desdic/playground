package evaluate

import (
	"log"
	"math"
	"testing"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		input  string
		output float64
		works  bool
	}{
		{"number", "1", 1, true},
		{"number fraction", "1.1", 1.1, true},
		{"add", "1+1", 2, true},
		{"add multiple", "1+1+1+1+1+1+1+1", 8, true},
		{"add fraction", "1.1+1", 2.1, true},
		{"sub", "1-1", 0, true},
		{"mul", "2*3", 6, true},
		{"div", "4/2", 2, true},
		{"addmul", "4+2*5", 14, true},
		{"subdiv", "6-4/2", 4, true},
		{"fraction", "2/3", 0.6666666666666666, true},
		{"fraction_neg", "-2/3", -0.6666666666666666, true},
		{"paren", "(1)", 1, true},
		{"paren neg", "-(1)", -1, true},
		{"paren calc", "(1+2)", 3, true},
		{"paren calc mul", "(1+2)*3", 9, true},
		{"unary minus", "-3", -3, true},
		{"unary minus", "1--3", 4, true},
		{"unassigned var", "bla", 0, false},
		{"sub sub", "-1-3", -4, true},
		{"abs", "abs(-3)", 3, true},
		{"abs expr", "abs(-3*2)", 6, true},
		{"floor", "floor(3.5)", 3, true},
		{"round", "round(3.5)", 4, true},
		{"AbC", "A*C*b", 42, true},
		{"AbC", "A+b+C", 24, true},
		{"AbC", "C+b+A", 24, true},
		// {"no operator", "2(3)", 6, false},
		// {"no operator", "2 6", 6, true},
	}

	vars := map[string]float64{
		"A": 21.0,
		"b": 1.0,
		"C": 2.0,
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			works := true
			f, err := Evaluate(tc.input, vars)
			if err != nil {
				works = false
			}

			if works != tc.works {
				log.Println(f)
				t.Fatal(err)
			}

			if !tc.works {
				return
			}
			res := math.Abs(f - tc.output)
			if res > 0.000001 {
				t.Fatal(f, "!=", tc.output)
			}
		})
	}
}

func benchmarkEvaluate(b *testing.B, str string, expect float64, vars map[string]float64) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		f, err := Evaluate(str, vars)
		if err != nil {
			log.Fatal(err)
		}
		res := math.Abs(f - expect)
		if res > 0.000001 {
			log.Fatal(f, "!=", expect)
		}
	}
}

func BenchmarkSimple(b *testing.B) { benchmarkEvaluate(b, "1+1", 2.0, nil) }
func BenchmarkVariables(b *testing.B) {
	benchmarkEvaluate(b, "A+b+C", 24.0, map[string]float64{"A": 21.0, "b": 1.0, "C": 2.0})
}

func BenchmarkComplex(b *testing.B) {
	benchmarkEvaluate(b, "44-(myvar-(-1*floor((4*abs(2*-1))/3)))", -2.0, map[string]float64{"myvar": 44.0})
}
