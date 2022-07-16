package main

import (
	"blog-service/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//new 一个Gin Engine 实例
	r := gin.New()
	// 初始化路由绑定
	bootstrap.SetupRoute(r)
	//运行服务，默认为 8080，我们指定端口为 8000
	err := r.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
