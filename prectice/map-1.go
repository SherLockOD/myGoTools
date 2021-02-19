package main

import "fmt"

// map 的值是指针
func main() {
	m := make(map[string]*string)
	a := "test"
	b := "1"
	m[a] = &b
	fmt.Println(m)

	c := "test"
	fmt.Println(*m[c])
}
