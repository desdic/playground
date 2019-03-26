// Package wstruct contains a single walk function for traversing a struct
package wstruct

import (
	"fmt"
	"reflect"
)

// WalkFunc is the definition of the Walk function
type WalkFunc func(aname string, avalue interface{}) error

// Walk traverses a struct
func Walk(input interface{}, fn WalkFunc) error {

	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Walk only work on structs")
	}

	vi := reflect.Indirect(v)
	for i := 0; i < v.NumField(); i++ {

		f := v.Field(i)
		if !f.CanInterface() {
			continue
		}

		if err := fn(vi.Type().Field(i).Name, f.Interface()); err != nil {
			return err
		}
	}
	return nil
}
