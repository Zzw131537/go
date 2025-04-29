package service

import (
	"HiChat/common"
	"HiChat/dao"
	"HiChat/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 新建群聊
func NewGroup(ctx *gin.Context) {
	owner := ctx.PostForm("ownerId")
	ownerId, err := strconv.Atoi(owner)
	if err != nil {
		zap.S().Info("owner 类型转换失败", err)
		return
	}

	ty := ctx.PostForm("cate")
	Type, err := strconv.Atoi(ty)
	if err != nil {
		zap.S().Info("ty类型转换失败", err)
		return
	}
	img := ctx.PostForm("icon")
	name := ctx.PostForm("name")
	desc := ctx.PostForm("desc")
	community := models.Community{}
	if ownerId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "您未登录",
		})
		return
	}
	if name == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "群名称不能为空",
		})
		return
	}
	if img != "" {
		community.Image = img
	}
	if desc != "" {
		community.Desc = desc
	}
	community.Name = name
	community.Type = Type
	community.OwnerId = uint(ownerId)

	code, err := dao.CreateCommunity(community)
	if err != nil {
		HandleErr(code, ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "建群成功",
	})
}

// 获取群列表
func GetGroupList(ctx *gin.Context) {
	owner := ctx.PostForm("ownerId")
	ownerId, err := strconv.Atoi(owner)
	if err != nil {
		zap.S().Info("owner类型转换失败", err)
		return
	}
	if ownerId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "您未登录",
		})
		return
	}

	rsp, err := dao.GetCommunityList(uint(ownerId))
	if err != nil {
		zap.S().Info("获取群列表失败", err)
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "你还没有加入任何群聊",
		})
		return
	}
	common.RespOKList(ctx.Writer, rsp, len(*rsp))
}

// 加入群聊
func JoinGroup(ctx *gin.Context) {
	comInfo := ctx.PostForm("comId")
	if comInfo == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "群名称不能为空",
		})
		return
	}

	user := ctx.PostForm("userId")
	userId, err := strconv.Atoi(user)
	if err != nil {
		zap.S().Info("user 类型转换失败")
		return
	}
	if userId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "你为登录",
		})
		return
	}
	code, err := dao.JoinCommunityByName(uint(userId), comInfo)
	if err != nil {
		HandleErr(code, ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "加群成功",
	})
}
