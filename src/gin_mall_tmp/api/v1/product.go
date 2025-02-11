/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 20:37:03
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//创建商品

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}

	// token := c.GetHeader("Authorization")

	// if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
	// 	token = token[7:]
	// }

	// claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
