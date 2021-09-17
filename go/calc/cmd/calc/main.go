package main

import (
	"fmt"

	"github.com/Desdic/playground/go/calc/internal/evaluate"
)

func main() {
	f, err := evaluate.Evaluate("(A+C+b)*2", map[string]float64{
		"A": 18.0,
		"b": 2.0,
		"C": 1.0,
	})
	if err != nil {
		fmt.Println("Error evaluating", err)

		return
	}

	fmt.Println("Result:", f)
}
