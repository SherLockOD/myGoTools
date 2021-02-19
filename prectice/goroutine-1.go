package main

import (
	"fmt"
	"sync"
)

// wg.Add 放在循环内外的情况，其实没区别
func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	fmt.Println(wg)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("I'm %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			fmt.Printf("I'm %d\n", i)
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("finish")
//}