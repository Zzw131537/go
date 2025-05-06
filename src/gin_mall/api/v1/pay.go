/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:40:06
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	oderPayservice := service.OrderPayService{}

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&oderPayservice); err == nil {
		res := oderPayservice.PayDown(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
