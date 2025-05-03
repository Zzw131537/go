package service

import (
	"HiChat/common"
	"HiChat/dao"
	"HiChat/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// user 对返回数据进行脱敏
type user struct {
	Name     string
	Avatar   string
	Gender   string
	Phone    string
	Email    string
	Identity string
}

// 好友列表
func FriendList(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Request.FormValue("userId"))
	users, err := dao.FriendList(uint(id))
	if err != nil {
		zap.S().Info("获取好友列表失败", err)
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "好友为空",
		})
		return
	}
	infos := make([]user, 0)
	for _, v := range *users {
		info := user{
			Name:     v.Name,
			Avatar:   v.Avatar,
			Gender:   v.Gender,
			Phone:    v.Phone,
			Email:    v.Email,
			Identity: v.Identity,
		}
		infos = append(infos, info)
	}
	common.RespOKList(ctx.Writer, infos, len(infos))
}

// 添加好友
// AddFriendByName 通过昵称添加好友

func AddFriendByName(ctx *gin.Context) {
	id := ctx.PostForm("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}
	tar := ctx.PostForm("targetName")
	target, err := strconv.Atoi(tar)
	if err != nil {
		code, err := dao.AddFriendByName(uint(userId), tar)
		if err != nil {
			HandleErr(code, ctx, err)
			return
		}
	} else {
		code, err := dao.AddFriend(uint(userId), uint(target))
		if err != nil {
			HandleErr(code, ctx, err)
			return
		}
	}
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "添加好友成功",
	})
}
func HandleErr(code int, ctx *gin.Context, err error) {
	switch code {
	case -1:
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	case 0:
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "该好友已存在",
		})
	case -2:
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "不能添加自己",
		})
	}
}

func RedisMsg(c *gin.Context) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	common.RespOKList(c.Writer, "ok", res)
}
