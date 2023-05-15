package main

import (
	"fmt"
)

// map 用来存储key-value的键值对

// 声明和初始化
func map1() {
	// 用map 表明 []内填上key类型 括号外填value类型·
	var m1 map[string]int //这样得到一个空的map 零值是nil 可以理解为空指针 不能直接操作m 否则会得到panic
	var m2 map[string]int
	// 需要进行初始化 使用make
	// e.g.
	m1 = make(map[string]int)
	//还可以指定创造出来的map的存储能力大小
	m2 = make(map[string]int, 100)
	// 也可以在声明的时候初始化
	var m3 = map[string]int{"abc": 3, "ccd": 4}
	// 也能通过赋值运算符 := 直接make出空的map来
	m4 := make(map[string]int)
	fmt.Println("m1:", m1, "\n m2:", m2, "\nm3:", m3, "\n m4:", m4)
}

// 增删改查
func map2() {
	//map 添加元素 直接用方括号赋值即可
	var m21 = make(map[string]int)
	//要保证m21经过初始化 否则会报nil的panic
	m21["acc"] = 4 //key值在map 已存在 会替换掉原本的key(更新) 不存在就是 添加
	m21["engine"] = 6
	fmt.Println(m21)

	// 删除元素 用delete  如果不存在就不会发生
	delete(m21, "acc") //要保证传入的map不为nil 否则会panic
	fmt.Println(m21)

	//查找元素 直接用[]进行查询
	// 如果key不再其中 只会返回error 所以用两个变量接收map的结果
	// 如果没有 error 就会得到nil 判断是否为nil 就知道元素是否存在了
	if val1, ok1 := m21["engine"]; ok1 {
		fmt.Println(val1)
	}

}
func main() {
	map1()
	fmt.Println("map2:")
	map2()
	// fmt.Println("==============")
	// wc.Test(WordCount)
}

// // 实战 map 例子
// func WordCount(s string) map[string]int {
// 	cnt := make(map[string]int)
// 	// 用Split方法拆分字符串
// 	for _, str := range strings.Split(s) {
// 		// 直接++ golang会自动填充
// 		cnt[str]++
// 	}
// 	return cnt
// }
