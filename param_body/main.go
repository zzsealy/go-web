package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context){

		bodyByts, err := ioutil.ReadAll(c.Request.Body)  // 读取数据流
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()  // 把输出结束
		}
		// 把获取到的body重新写入到请求内容中
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName := c.PostForm("first_name")
		// 默认值参数
		lastName := c.DefaultPostForm("last_name", "last_dafault_name")
		c.String(http.StatusOK, "%s, %s, %s", firstName, lastName, string(bodyByts))
	})
	 r.Run()
	 // 下面的返回值是如果不加22行的效果
	// curl -X POST "http://127.0.0.1:8080/test" -d "first_name=wang&last_name=quan"	    , last_dafault_name, first_name=wang&last_name=quan
	// 从上面的结果看出，没有获取到firstName 和 lastName 只是最后输出了整个body
}