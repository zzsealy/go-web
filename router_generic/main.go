package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	// get 泛绑定前缀
	r.GET("/User/*action", func(c *gin.Context){ 
	// 所有User前缀的都会返回这个路由 比如 curl -X GET "http://127.0.0.1:8080/User/dasd" 返回 hello world
		c.String(200, "hello world!")
	})
	
	r.Run()
}