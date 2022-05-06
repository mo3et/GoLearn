package main

func main() {
	//初始化
	var v2 int = 10 //写出变量类型并写出初始后的值
	var v3 = 10     //编译器通过值 自动匹配类型
	v4 := 10        //匹配类型然后赋值
	/**!!TODO! 同个变量不能被同时声明两次 下面两种为错误
	// var s string = "Hello"
	// var s string = "Hello"
	// s := "Hello"
	// s := "World"
	**/

	//变量赋值
	var a, b int
	a, b = b, a //交换两个变量

	// 支持匿名变量（对于不需要的返回值，不用额外定义去接收）

	// e.g. 函数返回两个变量，只取其一,可以这样写
	ret, _ := sample(a, b) // 已经定义变量 ret, _ =sample()
}

func sample(a, b int) (int, int) {
	return a, b
}
