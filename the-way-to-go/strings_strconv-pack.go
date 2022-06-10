// Articles: http://books.studygolang.com/the-way-to-go_ZH_CN/04.7.html

// 前缀和后缀
// HasPrefix 判断字符串 s 是否以 prefix 开头：

// strings.HasPrefix(s, prefix string) bool
// HasSuffix 判断字符串 s 是否以 suffix 结尾：

// strings.HasSuffix(s, suffix string) bool

// presuffix.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

// string contains
// Contains 判断string s 是否包含 substr

strings.Contains(s,substr string) bool 

// 判断子串 or 字符在父串出现的位置(索引)

// Index 返回string str 在string s中的索引 -1 表示s 不包含 str

strings.Index(s,str string) int

// LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引

strings.IndexRune(s string ,r rune) int
