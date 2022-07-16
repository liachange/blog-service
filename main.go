package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//new 一个Gin Engine 实例
	r := gin.New()
	//注册中间价
	r.Use(gin.Logger(), gin.Recovery())
	//注册路由
	r.GET("/ping", func(c *gin.Context) {
		//以JSON格式响应
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
	//运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}
