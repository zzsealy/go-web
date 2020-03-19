package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context){
		// 非默认值参数
		firstName := c.Query("first_name")
		// 默认值参数
		lastName := c.DefaultQuery("last_name", "last_dafault_name")

		c.String(http.StatusOK, "%s, %s", firstName, lastName)
	})

	r.Run(":8080")


	// curl -X GET "http://127.0.0.1:8080/test?first_name=wang"       wang, last_dafault_name%
	// curl -X GET "http://127.0.0.1:8080/test?first_name=wang&last_name=quan"   wang, quan
	
}