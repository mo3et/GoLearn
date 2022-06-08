func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/sumbit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	// Simple group :v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/sumbit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}
	router.Run(":8081")
}