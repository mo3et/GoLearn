package main

import (
	"fmt"
)

// Loop只有for  只支持i++ 不支持++i
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i := 0; ; i++ {
		if i > 10 {
			break
		}
	}
	for i := 0; i < 10; {
		i += 2
		fmt.Println(i)
	}
	fmt.Println("==================")
	ranges()
	fmt.Println("==================")
	mapstest1()
}

// Range usage
func ranges() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, v := range nums {
		sum += v
		fmt.Println(i, v)
	}
}

// map usage
func mapstest1() {
	kvs := map[string]string{"ap": "apple", "ba": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func sample() int {
	var v int
	return v
}

func sample2(a string) string {
	return a
}

//支持 if 与 switch的判断
func judges() {
	for i := 0; ; i++ {
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
	//在判断中初始化
	if v := sample(); v < 10 {
		fmt.Println(v)
	}

	//switch
	switch flag := sample2(); flag {
	//每个case 末尾不需要加break
	case "a":
		fmt.Println(flag)
	case "b":
		fmt.Println(flag)
	default:
		fmt.Println(flag)
	}

	//Switch的判断条件是按循序执行
	switch a := sample(); {
	case a < 5:
		fmt.Println(a)
	case a > 5:
		fmt.Println(a)
	default:
		fmt.Println("end")
	}
	// 上述代码 没有为switch设置判断的根据，逻辑等同于若干个if -else的罗列。在golang当中是同样允许的
}
