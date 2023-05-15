package main

import (
	"errors"
	"fmt"
)

// divide test
func Divide(a, b int) (ret int, err error) {
	if b == 0 {
		err = errors.New("divisor is zero")
		return
	}
	return a / b, nil
}

// 不定参数 用 ... 来定义  不定参数必须保证类型一样
func Sums1(nums ...int) int { //可以将int进行累加
	ret := 0
	for _, num := range nums {
		ret += num
	}
	return ret
}

// 等同于
func Sums2(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	return ret
}

// 不定参数使用方法
func arg() {
	a := make([]int)
	a = append(a, 3)
	a = append(a, 4)
	Sums2(a...)
	Sums2(a[1:]...)
}

func main() {
	ret, err := Divide(5, 2)
	if err == nil {
		fmt.Println(ret)
	} else {
		fmt.Println(err)
	}

	// 使用切片 传入四局必须强制转化
	Sums2([]int{3, 4, 6, 8})
	// 使用 ... 可以省略掉强制转化
	Sums1(3, 4, 5, 6)

	// Interface 用法是 可以用来代替所有类型的变量

	// 匿名函数
	sm := func(a, b int) int {
		return a + b
	}
	cc := sm(3, 4)

	//闭包 封闭外部
	af := func(x int) func(int) {
		return func(y int) {
			fmt.Println(x, y)
		}
	}
	// 真正封闭部分
	bf := af(4)
	bf(5)

}
