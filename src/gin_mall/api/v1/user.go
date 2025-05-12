/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 18:08:28
 */
package v1

import (
	"mall/pkg/util"
	"mall/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UserUpdate(c *gin.Context) {

	var userUpdate service.UserService
	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&userUpdate); err == nil {

		res := userUpdate.Update(c.Request.Context(), claims.ID)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UpLoadAvatar(c *gin.Context) {

	file, fileHead, _ := c.Request.FormFile("file")

	fileSize := fileHead.Size
	var upLoadAvatar service.UserService

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&upLoadAvatar); err == nil {

		res := upLoadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)
	if err := c.ShouldBind(&sendEmail); err == nil {

		res := sendEmail.Send(c.Request.Context(), claims.ID)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}

}

func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	if err := c.ShouldBind(&validEmail); err == nil {

		res := validEmail.Valid(c.Request.Context(), token)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}

}

func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService

	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	if err := c.ShouldBind(&showMoney); err == nil {

		res := showMoney.Show(c.Request.Context(), claims.ID)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}

}
