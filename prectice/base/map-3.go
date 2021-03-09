package main

import "fmt"

// S struct
type S struct {
	data string
}

func(s S) Read() string {
	return s.data
}

func(s *S) Write(str string) {
	s.data = str
}

// 使用值作为 receiver 的时候 method 可以通过指针调用，也可以通过值来调用。
func main()  {
	sVals := map[int]S{1: {"A"}}
	fmt.Println(sVals[1].Read())
	// sVals[1].Write("test") 不能编译通过

	sPtrs := map[int]*S{1: {"A"}}
	fmt.Println(sPtrs[1].Read())
	sPtrs[1].Write("test")
	fmt.Println(sPtrs[1].Read())
}

