package main

import "fmt"

type D struct {
	data string
}

func (d D) show() string {
	return d.data
}

func (d *D) change(str string) {
	d.data = str
}

func main() {
	// 相关 struct 不同值时，不用定义多个变量，使用map存储，并通过value调用method
	v := map[string]*D{
		"v1": {"A"},
		"v2": {"B"},
		"v3": {"C"},
	}

	// 循环show
	for k, v := range v {
		fmt.Printf("key: %s, value: %s\n", k, v.show())
	}

	// change 某个 key
	v["v3"].change("c")

	// show
	fmt.Println(v["v3"].show())

}
