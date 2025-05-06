/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-15 16:09:45
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateFavorite(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ListFavorite(c *gin.Context) {
	listFavoriteService := service.FavoriteService{}

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
