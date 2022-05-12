// This example will show how to serve static files like CSS, JavaScript or images from a specific directory.
// 展示如何从特定目录提供 static files

// static-files.go
package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

// $ tree assets/
// assets/
// └── css
//     └── styles.css

// Terminal
// $ go run static-files.go

// $ curl -s http://localhost:8080/static/css/styles.css
// body {
//     background-color: black;
// }
