package models

import (
	"HiChat/global"
	"errors"
)

type Community struct {
	Model
	Name    string // 群名称
	OwnerId uint   // 群拥有者
	Type    int    // 群类型
	Image   string // 群头像
	Desc    string // 描述
}

func (c *Community) ComTableName() string {
	return "community"
}

// 获取群成员id ,解决循环导包
func FindUsers(groupId uint) (*[]uint, error) {
	relation := make([]Relation, 0)
	if tx := global.DB.Where("target_id=?and type=2", groupId).Find(&relation); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到成员信息")
	}
	userIds := make([]uint, 0)
	for _, v := range relation {
		userId := v.OwnerId
		userIds = append(userIds, userId)
	}
	return &userIds, nil
}
