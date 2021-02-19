package main

import "fmt"

// 闭包
func temp(msg string) func() string {
	return func() string {
		msg = msg[:len(msg)-1]
		return msg
	}
}

func main() {
	t1 := temp("hello world")
	fmt.Println(t1())
	fmt.Println(t1())
	fmt.Println(t1())
}