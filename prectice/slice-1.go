package main

import "fmt"

type Data struct {
	data []string
}

func (d *Data) Set(list []string) {
	d.data = make([]string, len(list))
	copy(d.data, list)
}

func main() {
	d := Data{}
	list := []string{
		"a",
		"b",
		"c",
	}

	// 拷贝
	d.Set(list)

	// 结果展示
	fmt.Println(d.data)

	// 修改list，d.data 结果不变
	list[0] = "d"
	fmt.Println(d.data)
}

