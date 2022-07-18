package bootstrap

import (
	"blog-service/app/http/middlewares"
	"blog-service/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)
	//注册API路由
	routes.RegisterAPIRoutes(router)
	//配置404路由
	setup404Handler(router)
}
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
