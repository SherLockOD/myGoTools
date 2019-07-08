package base

import (
	"reflect"
)

func EqualMap(x, y interface{}) bool {

	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		return false
	}

	switch x.(type) {
	case map[string]string:
		x := x.(map[string]string)
		y := y.(map[string]string)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	case map[string]int:
		x := x.(map[string]int)
		y := y.(map[string]int)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	case map[int]int:
		x := x.(map[int]int)
		y := y.(map[int]int)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	case map[int]string:
		x := x.(map[int]string)
		y := y.(map[int]string)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	case map[string]bool:
		x := x.(map[string]bool)
		y := y.(map[string]bool)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	case map[int]bool:
		x := x.(map[int]bool)
		y := y.(map[int]bool)
		if len(y) != len(x) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || xv != yv {
				return false
			}
		}
	default:
		return false
	}
	return true
}
