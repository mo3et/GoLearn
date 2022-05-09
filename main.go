package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6}
	sum := 0
	for _, n := range a {
		sum += n
		fmt.Println(sum, n)
	}

	kvs := map[string]string{"a": "apple", "b": "ba", "c": "cheese", "d": "dd"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		// fmt.Println("key:", k, "value:", v)
	}

	// fmt.Println(a, b)
}
