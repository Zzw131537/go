package router

import (
	"HiChat/middlewear"
	"HiChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// 初始化路由
	router := gin.Default()

	// v1版本
	v1 := router.Group("v1")

	// 用户模块
	user := v1.Group("user")
	{
		user.GET("/list", middlewear.JWY(), service.List)
		user.POST("/login_pw", middlewear.JWY(), service.LoginByNameAndPassWord)
		user.POST("/new", middlewear.JWY(), service.NewUser)
		user.DELETE("/delete", middlewear.JWY(), service.DeleteUser)
		user.POST("/updata", middlewear.JWY(), service.UpdataUser)

	}

	relation := v1.Group("relation").Use(middlewear.JWY())
	{
		// 好友api
		relation.POST("/list", service.FriendList)
		relation.POST("/add", service.AddFriendByName)

		// 群api
		relation.POST("/new_group", service.NewGroup)
		relation.POST("/group_list", service.GetGroupList)
		relation.POST("/join_group", service.JoinGroup)
	}

	return router
}
