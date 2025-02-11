/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 19:01:46
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(da *gorm.DB) *ProductImgDao {
	return &ProductImgDao{da}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}
