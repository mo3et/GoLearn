// Website:https://gowebexamples.com/http-server/
package main

// TODO 重点查看 Handlerfunc  handle 对应的参数 是什么

// HTTP server 具备的功能
// 处理动态请求 处理来着浏览完整、登录用户或发布图像的用户的传入请求
// 提供静态资源：为浏览器提供JS、CSS和图像，为用户创造动态体验
// 接受连接：HTTPserver必须监听特定端口才能接收来自Internet的连接

//处理动态请求
// net/http包 包含接收请求并动态处理他们所需的handler
// 可以使用该函数注册一个新的handler http.HandlerFunc
// 他的第一个参数 需要一个匹配的路径 第二个参数是一个要执行的函数
http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"welcome to my website!")
})
// 对于动态方面，http.Request包含有关请求及其参数的所有信息。
//You can read GET parameters with   r.URL.Query().Get("token") 
// or POST parameters (fields from an HTML form) with    r.FormValue("email")


// 提供静态请求
// 为了提供JS、CSS、image等静态资源 使用内置的http.FileServer 并指向一个url路径
// 为了是file server 正常工作，他需要知道从哪提供文件
fs:=http.FileServer(http.Dir("static/"))

// 一旦服务器就位 只需要指向他的url 就像动态请求那样  
// 为了正确提供文件 需要去掉部分url 通常是我们文件所在目录的名称
http.Handle("/static/",http.StripPrefix("/static/",fs))

// 接受连接
http.ListenAndServe(":80", nil)

// Complete Code

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}