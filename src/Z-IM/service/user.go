package service

import (
	"HiChat/common"
	"HiChat/dao"
	"HiChat/middlewear"
	"HiChat/models"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户列表

func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "获取用户列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// 密码登录

func LoginByNameAndPassWord(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	data, err := dao.FindUserByName(name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "登录失败",
		})
		return
	}

	if data.Name == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不存在",
		})
		return
	}

	ok := common.CheckPassWord(password, data.Salt, data.PassWord)
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}

	Rsp, err := dao.FindUserByNameAndPwd(name, data.PassWord)
	if err != nil {
		zap.S().Info("登录失败", err)
	}

	// 使用jwt 做权限认证
	token, err := middlewear.GenerateToken(Rsp.ID, "yk")
	if err != nil {
		zap.S().Info("生成token失败", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"tokens":  token,
		"userId":  Rsp.ID,
	})
}

// 注册用户
func NewUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Request.FormValue("name")
	password := ctx.Request.FormValue("password")
	repassword := ctx.Request.FormValue("Identity")

	if user.Name == "" || password == "" || repassword == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空!",
			"data":    user,
		})
		return
	}

	// 查询用户是否存在
	_, err := dao.FindUser(user.Name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已注册",
			"data":    user,
		})
		return
	}

	if password != repassword {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "两次密码不一致!",
			"data":    user,
		})
		return
	}

	// 生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())

	// 加密密码
	user.PassWord = common.SaltPassWord(password, salt)
	user.Salt = salt
	t := time.Now()
	user.LoginTime = &t
	user.LoginOutTime = &t
	user.HeartBeatTime = &t
	dao.CreateUser(user)
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功",
		"data":    user,
	})
}

// 更新用户信息
func UpdataUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "类型转换失败",
		})
		return
	}
	user.ID = uint(id)
	Name := ctx.Request.FormValue("name")
	PassWord := ctx.Request.FormValue("password")
	Email := ctx.Request.FormValue("email")
	Phone := ctx.Request.FormValue("email")
	avatar := ctx.Request.FormValue("icon")
	gender := ctx.Request.FormValue("gender")
	if Name != "" {
		user.Name = Name
	}

	if PassWord != "" {
		salt := fmt.Sprintf("%d", rand.Int31())
		user.Salt = salt
		user.PassWord = common.SaltPassWord(PassWord, salt)
	}
	if Email != "" {
		user.Email = Email
	}
	if Phone != "" {
		user.Phone = Phone
	}
	if avatar != "" {
		user.Avatar = avatar
	}
	if gender != "" {
		user.Gender = gender
	}
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		zap.S().Info("参数不匹配", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数不匹配",
		})
		return
	}

	Rsp, err := dao.UpdateUser(user)
	if err != nil {
		zap.S().Info("更新用户失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "修改信息失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "修改成功",
		"data":    Rsp.Name,
	})
}

// 账号注销
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "类型转换失败",
		})
		return
	}
	user.ID = uint(id)
	err = dao.DeleteUser(user)
	if err != nil {
		zap.S().Info("注销账号失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "注销账号失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注销账号成功",
	})
}

// SendUserMsg 发送消息
func SendUserMsg(ctx *gin.Context) {
	models.Chat(ctx.Writer, ctx.Request)
}
