/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-16 16:08:43
 */
package route

import (
	"mall/middleware"
	"net/http"

	api "mall/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("product/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)
		authed := v1.Group("/") // 需要登录保护

		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UpLoadAvatar)

			// 更新密码需要邮箱验证

			authed.POST("user/sending-email", api.SendEmail)

			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金娥
			authed.POST("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)

			authed.POST("search_proudcts", api.SearchProduct)

			// 编写收藏夹操作
			authed.GET("favorites", api.ListFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorite/:id", api.DeleteFavorite)

			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.GET("listaddress", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)
		}
	}

	return r
}
