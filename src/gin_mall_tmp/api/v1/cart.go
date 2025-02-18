/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 18:23:52
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	createCartService := service.CartService{}
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ListCart(c *gin.Context) {
	listCartService := service.CartService{}

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&listCartService); err == nil {
		res := listCartService.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UpdateCart(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
