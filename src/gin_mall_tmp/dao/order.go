/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:33:26
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(da *gorm.DB) *OrderDao {
	return &OrderDao{da}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(oId, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id = ?", oId, uId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByCondition(t, uId uint, page model.BasePage) (orderes []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id = ? AND type = ?", uId, t).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&orderes).Error
	return
}

func (dao *OrderDao) DeleteOrder(uId, cId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=? AND user_id = ? ", cId, uId).Delete(&model.Order{}).Error
}
