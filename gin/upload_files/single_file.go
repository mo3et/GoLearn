func main() {
	router := gin.Default()
	// Set a lower memory limit ofr miltipart form (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8081")
}

// How to curl
// curl -X POST http://localhost:8080/upload \
// -F "file=@/Users/appleboy/test.zip" \
// -H "Content-Type: nultipart/form-data"