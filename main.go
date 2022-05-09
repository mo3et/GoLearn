package main

import (
	"fmt"
)

func main() {
	a := 4
	if a, b := 4; b {
		fmt.Println(a)
	}
	fmt.Println(a, b)
}
