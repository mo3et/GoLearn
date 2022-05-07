package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringtest1() {
	var str string
	str1 := "hello world"
	var str2 = "hello world"
	c := "heloo" + "peropero"
	fmt.Println(str)
	fmt.Println(str1[3])
	fmt.Println(str2)
	fmt.Println(len(str2))
	fmt.Println(c)
}

// UFT8编码需要注意的问题
func lenUTF8test1() {
	str1 := "hello 世界"
	fmt.Println(len(str1)) //设想返回8 结果是12 因为汉字需要3个字符编码
	//解决方法
	fmt.Println(len([]rune(str1))) //将String转化为rune数组 然后计算长度
}

//类型转换

// strconv 将字符转为其他类型
// String convert
func convertStr() {
	//转整数和浮点数 用ParseInt 和ParseUint 接受三个参数 "str" "进制" "bit大小"
	value, err := strconv.ParseInt("33225", 10, 32) //第一个为 要转类型的字符串 第二个为 字符串的进制 跌三个为 返回bit的大小
	if err != nil {
		fmt.Println("error happens")
	}
	fmt.Println(value)

	//简易版
	value2, err := strconv.Atoi("11234")
	if err != nil {
		fmt.Println("error happens")
	}
	fmt.Println(value2)

	// 转Float只用 ParseFloat 只有两个参数 str bit大小
	value3, err := strconv.ParseFloat("33.33", 32)
	if err != nil {
		fmt.Println("error happens")
	}
	fmt.Println(value3)
}

// Int Float convert
func convertnum() {
	num := 180
	fmt.Println(strconv.FormatInt(int64(num), 16))

	num1 := 1801
	fmt.Println(strconv.Itoa(num1)) //FormatInt(i,10)

	// Float转str 接受4个参数 "待转换Float" "希望转换得到的格式(f b 'e E' 'g G' )" "特殊精度 一般为 -1"
	num2 := 23423134.323422
	fmt.Println(strconv.FormatFloat(float64(num2), 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(float64(num2), 'b', -1, 64))
	fmt.Println(strconv.FormatFloat(float64(num2), 'e', -1, 64))
	fmt.Println(strconv.FormatFloat(float64(num2), 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(float64(num2), 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(float64(num2), 'G', -1, 64))
}

// str bool convert
func convertbool() {
	flag, err := strconv.ParseBool("t")
	if err != nil {
		fmt.Println("error happens")
	}
	fmt.Println(flag)
}

// 字符串运算包 strings
func stringstest1() {
	//**  Compare 比较字符串大小 返回int
	cmp := strings.Compare("a", "b")
	fmt.Println("Compare is:", cmp)

	//index 查找字符串中子串位置 会返回第一次出现的位置 否则返回-1
	var theInd = strings.Index("iufjdfhjsubasdlfj", "sub")
	// lastindex 找出出现的最后一个位置
	var theLastIndex = strings.LastIndex("adfjlastdflastlsat", "last")
	fmt.Println("The index is:", theInd, "\nand the last index is:", theLastIndex)

	//** 用Count统计子串在整体出现的次数
	var CountsNum = strings.Count("abc abc abc abc abc abc", "abc") //若子串为空 结果为 母串长度+1
	fmt.Println("The count is:", CountsNum)

	// ** 用Repeat 讲字符串重复指定的次数
	repeat := strings.Repeat("abc", 10)
	fmt.Println("The repeat is:", repeat)

	// Replace 替换字符串中的不服 接收四个参数 "str" "匹配串" "目标串" "替换次数"
	strRe := "aaaddc"
	strings.Replace(strRe, "a", "b", 1) //baaddc
	fmt.Println("First Replace:", strRe)
	strings.Replace(strRe, "a", "b", -1) // 如果替换次数 小于0 表示全部替换
	fmt.Println("Second Replace:", strRe)

	// ** Split 用来分隔字符串 传入两个参数 "str" "分隔符"
	strSp := "abc,bbc,bbd"
	sliceSp := strings.Split(strSp, ",")
	fmt.Println(sliceSp)

	// ** Join Split的逆操作 通过指定的分隔符 将字符串数组拼接
	sliceJo := []string{"aab", "aba", "baa"}
	strJo := strings.Join(sliceJo, ",")
	fmt.Println(strJo)
}

func main() {
	stringtest1()
	lenUTF8test1()
	convertStr()
	convertnum()
	convertbool()
	stringstest1()
}

//  总结
//  strconv是专门用来字符串和其他类型进行转换的
//  strings当中封装的是操作字符串的一些函数。比如字符串判断、join、split
//  等各种处理
