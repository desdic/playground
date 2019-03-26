package wstruct

import (
	"fmt"
	"reflect"
)

func ExampleWalk_first() {

	m := struct {
		Name string
		Age  int
	}{
		"Kim", 42,
	}

	err := Walk(m, func(aname string, avalue interface{}) error {
		switch avalue.(type) {
		case string:
			fmt.Println("Got the string:", aname, avalue)
		case int:
			fmt.Println("Got the int:", aname, avalue)
		default:
			return fmt.Errorf("Unhandled variable called %s, type %v", aname, reflect.TypeOf(avalue))
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Failed to walk struct: %v\n", err)
	}
	// Output:
	// Got the string: Name Kim
	// Got the int: Age 42
}

func ExampleWalk_second() {
	type mystruct struct {
		Name    string
		Age     int
		Parents *mystruct
		Weight  float32
	}

	m := mystruct{Name: "Kim", Age: 42}

	err := Walk(m, func(aname string, avalue interface{}) error {
		switch avalue.(type) {
		case string:
			fmt.Println("Got the string:", aname, avalue)
		case int:
			fmt.Println("Got the int:", aname, avalue)
		case *mystruct:
			fmt.Println("Got the *mystruct:", aname, avalue)
		default:
			return fmt.Errorf("Unhandled variable called %s, type %v", aname, reflect.TypeOf(avalue))
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Failed to walk struct: %v\n", err)
	}
	// Output:
	// Got the string: Name Kim
	// Got the int: Age 42
	// Got the *mystruct: Parents <nil>
	// Failed to walk struct: Unhandled variable called Weight, type float32
}

func ExampleWalk_third() {

	err := Walk(42, func(aname string, avalue interface{}) error {
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to walk struct: %v\n", err)
	}
	// Output: Failed to walk struct: Walk only work on structs
}
