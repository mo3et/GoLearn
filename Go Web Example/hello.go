package main

import (
	"fmt"
	"net/http"
)

// Hello World
// Website：https://gowebexamples.com/hello-world/

// Go Standard library net/http 包含所有http功能
// 包括HTTP client 和HTTP server
// 此example 将了解如何创建可以在浏览器中查看的网络服务器

// Registering a Request Handler
//创建一个来自browser HTTP clients or API requests 的handler
// 该func接收(receives) 两个参数(parameters)
// func (w http.ResponseWriter, r *http.Request)
// //ResponseWriter 编写text/html响应;
// // Request 包含有关此HTTP请求的所有信息 包括url or header fields等信息

// // 将handler注册到默认 HTTP server
// http.HandleFunc("/",dunc (w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w,"Hello,you've request: %s\n",r.URL.Path)
// })

// Listen for HTTP Connections
// 单独的handler不能接受来自外部的HTTP连接
// HTTP server必须 listen(监听) on a port 才能将连接传递给请求 handler
// 因为 port 80 大多情况下是default port for HTTP traffic 所以server will also listen on it

// 启动默认的HTTPserver 并 listen for connections on port 80(监听80端口上的连接)
// 可以在localhost, see your server handing your request(查看服务器正在处理您的请求)
// http.ListenAndServe(":80", nil)

// complete code
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":8087", nil)
}
