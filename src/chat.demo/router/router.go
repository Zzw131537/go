package router

import "github.com/gin-gonic/gin"

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
	}
	return r
}
