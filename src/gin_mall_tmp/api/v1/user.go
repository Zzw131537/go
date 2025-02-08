/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-07 19:30:19
 */
package v1

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, err)
	}
}
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(c *gin.Context) {

	var userUpdate service.UserService
	token := c.GetHeader("Authorization")

	if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
		token = token[7:]
	}

	claims, _ := util.ParseToken(token)

	fmt.Println("进入api_Update函数")
	if err := c.ShouldBind(&userUpdate); err == nil {
		fmt.Println("进入CC")

		res := userUpdate.Update(c.Request.Context(), claims.ID)

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
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
		c.JSON(http.StatusBadRequest, err)
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
		c.JSON(http.StatusBadRequest, err)
	}

}
