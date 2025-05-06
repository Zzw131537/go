package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"

	"go.uber.org/zap"
)

// FriendList 获取好友列表
func FriendList(userId uint) (*[]models.UserBasic, error) {
	relation := make([]models.Relation, 0)
	if tx := global.DB.Where("owner_id = ? and type =1", userId).Find(&relation); tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation数据")
		return nil, errors.New("未查询到好友关系")
	}

	userID := make([]uint, 0)
	for _, v := range relation {
		userID = append(userID, v.TargetID)
	}
	user := make([]models.UserBasic, 0)
	if tx := global.DB.Where("id in ?", userID).Find(&user); tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation好友关系")
		return nil, errors.New("未查询到好友")
	}
	return &user, nil
}

// 通过id 添加好友 需要进行事务管理
func AddFriend(userId, targetId uint) (int, error) {
	if userId == targetId {
		return -2, errors.New("userId 与targetId 相等")
	}

	// 通过id 查询用户
	targetUser, err := FindUserID(targetId)
	if err != nil {
		return -1, errors.New("未查询到用户")
	}
	if targetUser.ID == 0 {
		zap.S().Info("未查询到用户")
		return -1, errors.New("未查询到用户")
	}
	relation := models.Relation{}
	if tx := global.DB.Where("owner_id = ? and target_id =? and type =1", userId, targetId).First(&relation); tx.RowsAffected == 1 {
		zap.S().Info("该好友已存在")
		return 0, errors.New("该好友已存在")
	}

	if tx := global.DB.Where("owner_id = ? and target_id =?and type =1", targetId, userId).First(&relation); tx.RowsAffected == 1 {
		zap.S().Info("该好友已存在")
		return 0, errors.New("该好友已存在")
	}

	// 开启事务
	tx := global.DB.Begin()
	relation.OwnerId = userId
	relation.TargetID = targetUser.ID
	relation.Type = 1
	if t := tx.Create(&relation); t.RowsAffected == 0 {
		zap.S().Info("创建失败")
		// 事务回滚
		tx.Rollback()
		return -1, errors.New("创建好友记录失败")
	}
	relation = models.Relation{}
	relation.OwnerId = targetId
	relation.TargetID = userId
	relation.Type = 1
	if t := tx.Create(&relation); t.RowsAffected == 0 {
		zap.S().Info("创建失败")

		// 事务回滚
		tx.Rollback()
		return -1, errors.New("创建好友记录失败")
	}

	// 提交事务
	tx.Commit()
	return 1, nil
}

// 通过昵称进行添加

func AddFriendByName(userId uint, targetName string) (int, error) {
	user, err := FindUserByName(targetName)
	if err != nil {
		return -1, errors.New("该用户不存在")
	}
	if user.ID == 0 {
		zap.S().Info("未查询到该用户")
		return -1, errors.New("该用户不存在")
	}
	return AddFriend(userId, user.ID)
}
