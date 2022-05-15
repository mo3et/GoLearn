package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var a int = 0
	var b int = 1
	// sum := 0
	return func() int {
		// sum = a + b
		// a = b
		// b = sum
		// return a
		//return sum
		c := a
		a = b
		b = a + c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
