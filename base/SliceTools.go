package base

import (
	"fmt"
	"reflect"
)

func AllIntEqual(s []int) (bool, error) {
	l1 := len(s)
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	l2 := len(m)
	if l1 < 2 {
		return false, fmt.Errorf("too few elements %v", s)
	}
	if l2 != 1 {
		return false, fmt.Errorf("all of the nums '%v' are not equal", s)
	}
	return true, nil
}

func EqualSlice(x, y interface{}) (bool, error) {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		return false, fmt.Errorf("slices type is not equal")
	}

	switch x.(type) {
	case []int:
		x := x.([]int)
		y := y.([]int)
		if len(x) != len(y) {
			return false, fmt.Errorf("slices length is not equal")
		}
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				return false, fmt.Errorf("slices value '%v' and '%v' is not equal", x[i], y[i])
			}
		}
	case []string:
		x := x.([]string)
		y := y.([]string)
		if len(x) != len(y) {
			return false, fmt.Errorf("slices length is not equal")
		}
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				return false, fmt.Errorf("slices value '%v' and '%v' is not equal", x[i], y[i])
			}
		}
	case []bool:
		x := x.([]bool)
		y := y.([]bool)
		if len(x) != len(y) {
			return false, fmt.Errorf("slices length is not equal")
		}
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				return false, fmt.Errorf("slices value '%v' and '%v' is not equal", x[i], y[i])
			}
		}
	default:
		return false, fmt.Errorf("slices type is not support")
	}
	return true, nil
}
