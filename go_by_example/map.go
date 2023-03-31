// Map 是 Go 内建的关联数据类型 (hash or dict)
package main

import "fmt"

func main() {
	// 创建空 Map 需要使用内建函数 make:
	// make(map[key-type]val-type) //key value
	m := make(map[string]int)

	// 使用 name[key]= val 语法来设置 键值对
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// 使用 name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len 函数 可以返回 map的 key-value 数量
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个map中
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
