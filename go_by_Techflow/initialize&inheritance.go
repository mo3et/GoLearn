package main

import (
	"fmt"
)

/*
	struct initialize 有四种方法
	"new关键字" "结构体名称" "" ""
*/

// new
type Point1 struct {
	x int
	y int
}

func main() {
	var p1 *Point1 = new(Point1)
	fmt.Print(p1)
}
