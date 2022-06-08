// Redirect (重定向)

// Issuing a HTTP redirect.
// Both internal and external locations are supported.

r.GET("/test", func(c *gin.Context) {
	c.Redirect(htpp.StatusMovedPermanently, "http://www.google.com/")
})

// 通过 POST 进行 HTTP redirect
r.POST("/test", func(c *gin.Context) {
	c.Redirect(http.StatusFound, "/foo")
})

// Routes Redirect, Use HandleContext:
r.GET("/test", func(c *gin.Context) {
	c.Request.URL.Path = "/test2"
	r.HandleContext(c)
})
r.GET("/test2",func(c *gin.Context) {
	c.JSON(200,gin.H{"hello:"world})
})