package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("testing", testing)
	r.Run()
}
// curl -X GET "http://127.0.0.1:8080/testing?name=wang&address=beijing&birthday=2000-01-02"                 {wang beijing 2000-01-02 00:00:00 +0800 CST}
// curl -X POST "http://127.0.0.1:8080/testing" -d 'name=wang&address=beijing&birthday=2000-01-02'           {wang beijing 2000-01-02 00:00:00 +0800 CST}
// curl -H "Content-Type:application/json" -X POST 'http://127.0.0.1:8080/testing' -d '{"name": "wang"}'     {wang  0001-01-01 00:00:00 +0000 UTC}

func testing(c *gin.Context) {
	var person Person
	// 这里根据请求的content-type来做不同的binding操作
	if err := c.ShouldBind(&person); err == nil  { // 如果没有错误
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error:%v", err)
	}

}
