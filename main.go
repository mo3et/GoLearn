package main

import (
	"fmt"
)

// func main() {
// 	a := []int{2, 4, 6}
// 	sum := 0
// 	for _, n := range a {
// 		sum += n
// 		fmt.Println(sum, n)
// 	}

// 	kvs := map[string]string{"a": "apple", "b": "ba", "c": "cheese", "d": "dd"}
// 	for k, v := range kvs {
// 		fmt.Printf("%s -> %s\n", k, v)
// 		// fmt.Println("key:", k, "value:", v)
// 	}

// 	// fmt.Println(a, b)
// }

//  Channel 和 Goroutine 实现
// func Add(x, y int, c chan int) {
// 	z := x + y
// 	c <- z

// }

// func main() {
// 	c := make(chan int)
// 	go Add(3, 4, c)
// 	z := <-c
// 	fmt.Println(z)

// }

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
