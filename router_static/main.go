package main

import "github.com/gin-gonic/gin"
import "net/http"

func main(){
	r := gin.Default()
	r.Static("/assets", "./assets")
	// 另一种设置方法
	r.StaticFS("/static", http.Dir("static"))
	// 设置单个静态文件
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Run()

}