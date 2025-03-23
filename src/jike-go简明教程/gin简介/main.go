/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-22 21:24:29
 */
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-21 16:05:15
 */

/*
特性:
快速:路由不使用反射，基于Redix树,内存占用较少
中间件:
异常处理
JSON
路由分组
渲染内置
*/

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello ,Geektutu")
	})

	// 解析路径参数
	r.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	// 获取Query 参数
	r.GET("/user", func(ctx *gin.Context) {
		name := ctx.Query("name")
		role := ctx.DefaultQuery("role", "teacher")
		ctx.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 获取Post 参数
	r.POST("/form", func(ctx *gin.Context) {
		username := ctx.PostForm("username") // 获取formData 中的参数
		password := ctx.DefaultPostForm("password", "000000")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// Query 与Post混合参数
	r.POST("/posts", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password", "00000")

		ctx.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// 重定向
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	})
	r.GET("/goindex", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/"
		r.HandleContext(ctx)
	})

	// 分组路由
	defaultHandler := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"path": ctx.FullPath(),
		})
	}

	// group v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("series", defaultHandler)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("posts")
		v2.GET("series")
	}

	// 上传文件
	r.POST("upload1", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		ctx.String(http.StatusOK, "%s is upload", file.Filename)
	})

	// 多个文件
	r.POST("/upload2", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
		}
		ctx.String(http.StatusOK, "%d files upload", len(files))
	})

	r.GET("/product", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "商品")
	})
	// 作用于全局
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// // 作用于单个路由
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// // 作用于某个组
	// authorized := r.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// }

	// 热加载调试 Hot Reload
	/*
			github.com/codegangsta/gin
		github.com/pilu/fresh
	*/

	// go get -v -u github.com/pilu/fresh

	/*
				$ fresh
				fresh：未找到命令

		go install github.com/pilu/fresh@latest
	*/
	//安装好后，只需要将go run main.go命令换成fresh即可。每次更改源文件，代码将自动重新编译(Auto Compile)。
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
