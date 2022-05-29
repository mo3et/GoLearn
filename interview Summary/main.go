package main

import "fmt"

// func arrs() []int {
// 	s := make([]int, 0)
// 	for i := 0; i < 10; i++ {
// 		s = append(s, i)
// 	}
// 	return s
// }
// func main() {
// 	fmt.Println(arrs())
// }

// TODO
// 将数字转为2进制，然后计算二进制中的1的个数
func exam1(num int) {
	// var num1 int
	sum := 0
	fmt.Scanf("%d", &num)
	for num < 0 {
		if num%2 != 0 {
			sum += 1
			// num=num/2
		}
		num = num / 2
	}
	fmt.Println(sum)
}
func main() {
	exam1(9)
}

//奇偶排序数列 只能用一个slice

// 两个字符串 找其中的共同的子串
