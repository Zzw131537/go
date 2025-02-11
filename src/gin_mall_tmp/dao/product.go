/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 20:27:23
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

func (dao *ProductDao) CountProductByConfition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}
