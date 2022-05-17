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
