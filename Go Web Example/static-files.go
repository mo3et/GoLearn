// Assets and Files
// how to serve static files from a specific directory (特定目录)

// static-files.go
package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8081", nil)
}

// $ tree assets/
// assets/
// └── css
// └── styles.css
//
