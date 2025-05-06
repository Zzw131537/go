/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 21:38:31
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(da *gorm.DB) *CategoryDao {
	return &CategoryDao{da}
}

func (dao *CategoryDao) ListCategory() (category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return

}
