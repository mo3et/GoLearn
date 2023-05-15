package main

import "fmt"

// 自定义类型
// Go可以使用type关键字来自定义类型

// 可以基于内置的基本类型定义 也可以通过struct定义 e.g.
// 将MyInt定义为int类型
type MyInt1 int

// 类型别名
// TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型
type TypeAlias = Type //Type

// rune和byte就是类型别名 e.g.
type byte = uint8
type rune = uint32

// 类型定义和类型别名
// 类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a: main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b: int
}
