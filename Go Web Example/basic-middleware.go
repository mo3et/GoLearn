// basic-middleware.go

// e.g. 在Go中创建基本的 logging middleware (日志中间件)
// 中间件只需将 htto.HandleFunc 作为其参数之一，将其包装并返回一个新的 http.HandleFunc供 server调用

package mian

import {
	"fmt"
	"log"
	"net/http"
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		
	}
}