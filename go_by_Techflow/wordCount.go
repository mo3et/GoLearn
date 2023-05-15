// 实战 map 例子
package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	cnt := make(map[string]int)
	// 用Split方法拆分字符串
	for _, str := range strings.Split("abacdcs,fadsf", ",") {
		// 直接++ golang会自动填充
		cnt[str]++
	}
	return cnt
}
func main() {
	wc.Test(WordCount)
	fmt.Println("")
}
