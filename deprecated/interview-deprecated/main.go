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

func slice1() {
	data := []int{1, 2, 3}
	for _, v := range data {
		// v*=10 //data中原有元素是不会被修改的
		data[v] = v * 10
	}
	fmt.Println("data :", data)
}

func slice2() {
	data2 := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data2 {
		v.num *= 10 // 直接使用指针更新
	}
	fmt.Println(data2[0], data2[1], data2[2])
}

// TODO
// 将数字转为2进制，然后计算二进制中的1的个数
func exam1(num int) {
	// var num1 int
	sum := 0
	// fmt.Scanf("%d", &num)
	for num > 0 {
		if num%2 != 0 {
			sum += 1
			// num=num/2

		}
		num = num / 2
		fmt.Println(num, sum)
	}
	fmt.Println("sum", sum)
}
func main() {
	exam1(54)
	slice2()
}

//奇偶排序数列 只能用一个slice

// 两个字符串 找其中的共同的子串
