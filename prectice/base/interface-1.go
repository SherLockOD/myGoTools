package main

import "fmt"

type Mobile interface {
	call()
	sales() int
}

type Nokia struct {
	price int
}

type Iphone struct {
	price int
}

func (n Nokia) call() {
	fmt.Println("Nokia is me.")
}

func (n Nokia) sales() int {
	return n.price
}

func (i Iphone) call() {
	fmt.Println("Iphone is me.")
}

func (i Iphone) sales() int {
	return i.price
}

func main() {
	//var mobile = [5]Mobile{
	//	Nokia{350},
	//	Iphone{5000},
	//	Iphone{3400},
	//	Nokia{450},
	//	Iphone{5000},
	//}
	var mobile = [5]Mobile{
		Nokia{},
		Iphone{},
		Iphone{},
		Nokia{},
		Iphone{},
	}

	var totalSales = 0
	for _, m := range mobile {
		m.call()
		totalSales += m.sales()
	}
	fmt.Println(totalSales)
	Nokia{200}.call()
	Iphone{}.call()
}