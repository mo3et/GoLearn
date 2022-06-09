package main

import (
	"fmt"
)

// func main() {
//     // 合起来写
//     go func() {
//         i := 0
//         for {
//             i++
//             fmt.Printf("new goroutine: i = %d\n", i)
//             time.Sleep(time.Second)
//         }
//     }()
//     i := 0
//     for {
//         i++
//         fmt.Printf("main goroutine: i = %d\n", i)
//         time.Sleep(time.Second)
//         if i == 20 {
//             break
//         }
//     }
// }

// 用goroutine去接收 channal value
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) //启用go Routine
	ch <- 10
	// num := <-ch
	// select {
	// case <-ch:
	// 	// ch <- 10
	// 	fmt.Println("Great.", num)
	// }
	fmt.Println("Send Success.")
}
