/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-21 19:23:31
 */
package router

import (
	"chat/api"
	"chat/service"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	// Recover
	// Log 日志
	v1 := r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.GET("ws", service.Handler)
	}
	return r
}
