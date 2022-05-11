// Website:https://gowebexamples.com/http-server/
package main

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
// 对于动态方面，http.Request包含有关请求及其参数的所有信息。可以