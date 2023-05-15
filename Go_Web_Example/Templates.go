package main

// html/template包为HTML模板提供了丰富的模板语言 主要用于web application在client browser中以结构化方式显示数据
// Go模板语言好处是 自动转义 无需担心xss攻击
//  Go 会解析 HTML 模板并在将其显示给浏览器之前 转义所有输入。

// First Templates
// e.g. TODO List example
// 以HTML中的无序列表 ul 形式编写
data:= TodoPageData{
	PageTitle: "My TODO List",
	Todos:[]Todo{
		{Title:"Task 1",Done:false}
		{Title:"Task 2",Done:true}
		{Title:"Task 3",Done:true}
	}
}

// html Templates
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>

// 控制结构
// 模板语言包含一组丰富的控制结构来呈现您的 HTML。在这里，您将获得最常用的概述。

// Control Structure				Definition
// {{/* a comment */}}				Defines a comment
// {{.}}							Renders the root element
// {{.Title}}						Renders the “Title”-field in a nested element
// {{if .Done}} {{else}} {{end}}	Defines an if-Statement
// {{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
// {{block "content" .}} {{end}}	Defines a block with the name “content”


// 模板可以从string或 a file on Disk 
// e.g. 在Go 程序的同一目录 有layout.html的模板文件

tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))

// 在Request Handler 中执行 Templates 
// 从磁盘解析模板，他就可以在request Handler中使用，
// Execute 函数接收一个 io.Writer来写出模板 和 一个interface{} 来将数据传递到模板上

// 在http.ResponseWriter上调用该函数时, ContentType 会自动设置在对应ContentType的HTTP Response  
// Content-Type: text/html; charset=utf-8

func (w http.ResponseWriter,r *http.Request) {
	tmpl.Execute(w,"data goes here")
}

// Complete Code
package main


import (
    "html/template"
    "net/http"
)



type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}



func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":80", nil)
}

// HTML templates
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>