package clousure

func closures(){

    // Closures

    // 在闭包中,匿名函数 add 返回了一个匿名函数,
    // 该匿名函数可以反问外部函数 add 的变量x.
    // 通过调用add(2) 创建了一个闭包函数 addTwo
    // 然后调用addTwo(3), 会将3 与外部变量x(值为2) 相加,得到结果5.
    add:=func(x int) func(int) int{
        return func(y int) int{
            return x+y
        }
    }
    addTwo:=add(2)
    fmt.Println(addTwo(3)) //Output 5
}
