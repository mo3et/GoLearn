// Gin allow by default use only one html.Template. Check a multitemplate render for using features like go 1.6 block template.

// Multitemplate renders
// Link: https://github.com/gin-contrib/multitemplate

// Simple example
package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})
	router.Run(":8080")
}
