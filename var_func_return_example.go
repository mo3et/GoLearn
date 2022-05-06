package main

import "fmt"

// 使用包装函数
func main() {

	_, a := num(1, 2)

	prn(a)

	prn1(num(1, 2))

}

func num(a, b int) (int, int) {

	return a + b, a - b

}

func prn(a int) {

	fmt.Println(a)

}

func prn1(_, b int) {

	prn(b)

}
