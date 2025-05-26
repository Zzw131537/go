package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定义一个中间件
func myHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("usersession", "userid-1")

		ctx.Next() // 放行
		ctx.Abort()
	}
}
func main() {
	ginServer := gin.Default()

	ginServer.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg":   "200",
			"value": "OK",
		})
	})
	userServer := ginServer.Group("/user")
	userServer.Use(myHandler())
	{
		userServer.GET("/info", func(ctx *gin.Context) {
			usersession := ctx.MustGet("usersession")
			log.Println(usersession)
			fmt.Println(usersession)
			ctx.JSON(200, gin.H{
				"code": "200",
				"msg":  "详细",
			})
		})
	}
	ginServer.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 拦截器
	// 中间件

	ginServer.Run(":8080")
}
