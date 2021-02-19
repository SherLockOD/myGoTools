package main

import "fmt"

// map 是引用类型，作为形参时，自带指针效果

func change(m map[string]string) {
	m["test"] = "2"
}

func main() {
	m := make(map[string]string)
	m["test"] = "1"
	fmt.Println(m)

	change(m)
	fmt.Println(m) // map[test:2] , map自带指针
}
