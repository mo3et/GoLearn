package main

r:=mux.NewRouter()

// 注册Handler
// 是像调用 HandleFunc那样  r.HandleFunc()

// url 参数 
// mux最大优势是能从请求url 中提取段 e.g. 
// /books/go-programming-blueprint/page/10

// 此URL有两个动态段: 
// 书名slug(go-programing-blueprint)  page(10)

// 要让Handler与上面的URL匹配 可以使用占位符替换动态段
r.HandleFunc("/books/{titles}/page/{page",func(w http.ResponseWriter, r *http.Request){
	// get the books
	// navigate to the page
})

// 最后从这些段中获取数据，该包附带函数  mux.Vars(r) 他以http.request作为参数并返回段的映射
func(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	vars["titles"] //the book title slug
	vars["page"] //the page
	} 

	// 设置HTTP server's Router
为什么监听http.ListenAndServe(":80", nil)里面是nil,因为是HTTP server 的 Main Router 参数. 默认为nil
// nil 意味为使用 net/http默认路由器 要使用自己的路由器 nil ->r
http.ListnAndServe(":80",r)


	// Complete Code

	package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}

// Features of the gorilla/mux Router

// Methods
// Restrict the request handler to specific HTTP methods.
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// Hostnames & Subdomains
// Restrict the request handler to specific hostnames or subdomains.
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// Schemes
// Restrict the request handler to http/https.
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

// Path Prefixes & Subrouters
// Restrict the request handler to specific path prefixes.
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)