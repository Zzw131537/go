package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
)

// 获取群成员id
func FindUsers(groupId uint) (*[]uint, error) {
	relation := make([]models.Relation, 0)
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

// 新建群
func CreateCommunity(community models.Community) (int, error) {
	com := models.Community{}
	// 查询群是否已存在
	if tx := global.DB.Where("name = ?", community.Name).First(&com); tx.RowsAffected == 1 {
		return -1, errors.New("当前群记录已存在")
	}
	tx := global.DB.Begin()
	if t := tx.Create(&community); t.RowsAffected == 0 {
		t.Rollback()
		return -1, errors.New("群记录创建失败")
	}
	relation := models.Relation{}
	relation.OwnerId = community.OwnerId
	relation.TargetID = community.ID
	relation.Type = 2
	if t := tx.Create(&relation); t.RowsAffected == 0 {
		tx.Rollback()
		return -1, errors.New("群记录创建失败")
	}
	tx.Commit()
	return 0, nil
}

// 获取群列表
func GetCommunityList(ownerId uint) (*[]models.Community, error) {
	// 获取我加入的群
	relation := make([]models.Relation, 0)
	if tx := global.DB.Where("owner_id=? and type = 2", ownerId).Find(&relation); tx.RowsAffected == 0 {
		return nil, errors.New("不存在群记录")
	}

	communityID := make([]uint, 0)
	for _, v := range relation {
		cid := v.TargetID
		communityID = append(communityID, cid)
	}
	community := make([]models.Community, 0)
	if tx := global.DB.Where("id in ?", communityID).Find(&community); tx.RowsAffected == 0 {
		return nil, errors.New("获取数据失败")
	}
	return &community, nil
}

// 根据群昵称搜索并加入群
func JoinCommunityByName(ownerId uint, cname string) (int, error) {
	community := models.Community{}
	if tx := global.DB.Where("name = ?", cname).First(&community); tx.RowsAffected == 0 {
		return -1, errors.New("群记录不存在")
	}

	// 重复加群
	relation := models.Relation{}
	if tx := global.DB.Where("owner_id = ? and target_id = ? and type =2", ownerId, community.ID).First(&relation); tx.RowsAffected == 1 {
		return -1, errors.New("该群已经加入")
	}
	relation = models.Relation{}
	relation.OwnerId = ownerId
	relation.TargetID = community.ID
	relation.Type = 2
	if tx := global.DB.Create(&relation); tx.RowsAffected == 0 {
		return -1, errors.New("加入失败")
	}
	return 0, nil
}
