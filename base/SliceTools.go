package base

import (
	"errors"
	"fmt"
	"reflect"
)

type sliceHelper struct {
	slicePtr interface{}
}

// 参数slicePtr必须是指向slice的指针
func SliceHelper(slicePtr interface{}) *sliceHelper {
	return &sliceHelper{slicePtr}
}

// 只能用于移除slice的数组中的元素
func (t *sliceHelper) Remove(index int) error {
	if t.slicePtr == nil {
		return errors.New("slice ptr is nil")
	}

	slicePtrValue := reflect.ValueOf(t.slicePtr)
	// 必须为指针
	if slicePtrValue.Type().Kind() != reflect.Ptr {
		return errors.New("should be slice ptr")
	}

	sliceValue := slicePtrValue.Elem()
	// 必须为slice
	if sliceValue.Type().Kind() != reflect.Slice {
		return errors.New("should be slice ptr")
	}

	if index < 0 || index >= sliceValue.Len() {
		return errors.New("index out of range")
	}
	sliceValue.Set(reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, sliceValue.Len())))
	return nil
}


//
func SliceFilter(slice []string, elem string) []string {
	TS := make([]string, 0, len(slice) - len(elem))
	for _, v := range slice {
		if v == elem {
			continue
		}
		TS = append(TS, v)
	}
	return TS
}

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


func IsExists(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
