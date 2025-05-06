/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 21:07:20
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	listProductImgService := service.ProductImgService{}
	if err := c.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
