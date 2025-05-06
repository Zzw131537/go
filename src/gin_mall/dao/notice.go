/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-07 19:13:12
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewBiticeDaoByDB(da *gorm.DB) *NoticeDao {
	return &NoticeDao{da}
}

// GetUserById 根据id 获取Notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
