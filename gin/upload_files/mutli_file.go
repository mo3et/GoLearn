func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default  is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 //8 MiB
	router.POST("/upload", func(c *gin.Context) {
		//  Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			logPrintln(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8081")
}

// How to curl:
// curl -X POST http://localhost:8080/upload \
//   -F "upload[]=@/Users/appleboy/test1.zip" \
//   -F "upload[]=@/Users/appleboy/test2.zip" \
//   -H "Content-Type: multipart/form-data"