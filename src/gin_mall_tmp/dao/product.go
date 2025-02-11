/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 18:51:28
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(da *gorm.DB) *ProductDao {
	return &ProductDao{da}
}

func (dao *ProductDao) CreateProduct(product *model.Product) (er error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
