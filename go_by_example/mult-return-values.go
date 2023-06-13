// Go 原生支持多返回值。 这个特性很常用，例如用来 同时返回一个函数的结果和错误信息。
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// func vals1(a, b int) (int, int) {

// 	return a * b, a / b
// }
// func text(a int) int {
// 	c := b(int) - a
// 	if _, ok := text(c); ok {
// 		return 1
// 	}
// }

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	// m := make(map[int]int)

	// fmt.Println(error)

	// if _, ok := m[1]; ok {
	// 	println("ok", a)
	// }
	// _,ok:=vals1(6,3);ok{
	// 	reutrn nil;
	// }
}
