package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/get", func (c *gin.Context){
		c.String(200, "get方法")
	})

	r.POST("/post", func (c *gin.Context){
		c.String(200, "post方法")
	})

	r.Handle("DELETE", "/delete", func (c *gin.Context){
		c.String(200, "delete方法")
	}) 
	r.Any("/any", func( c *gin.Context){
		c.String(200, "any")
	})

    r.Run()

}