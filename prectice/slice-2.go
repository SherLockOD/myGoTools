package main

import "fmt"

type Data struct {
	data []string
}

func (d *Data) Get() []string {
	tmp := make([]string, len(d.data))
	copy(tmp, d.data)
	return tmp
}

func main() {
	data := Data{
		[]string{
			"1",
			"2",
		},
	}

	// 赋值
	value := data.Get()

	// 结果展示
	fmt.Println(value) // [1 2]

	// struct 修改值
	data.data[0] = "a"
	fmt.Println(data.data) // [a 2]
	fmt.Println(value) // [1 2]
}

