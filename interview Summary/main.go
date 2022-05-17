package main

import "fmt"

func arrs() []int {
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	return s
}
func main() {
	fmt.Println(arrs())
}

// TODO
// 将数字转为2进制，然后计算二进制中的1的个数

//奇偶排序数列 只能用一个slice

// 两个字符串 找其中的共同的子串
