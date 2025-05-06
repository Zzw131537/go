/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 19:56:46
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(da *gorm.DB) *CartDao {
	return &CartDao{da}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByaId(aId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", aId).First(&cart).Error
	return
}

func (dao *CartDao) ListCartByuId(uId uint) (cartes []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&cartes).Error
	return
}

func (dao *CartDao) UpdateCartByUserId(cart *model.Cart, uId, cId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id =?", cId, uId).Updates(&cart).Error
}

func (dao *CartDao) DeleteCart(uId, cId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=? AND user_id = ? ", cId, uId).Delete(&model.Cart{}).Error
}

func (dao *CartDao) UpdateCartNumBuUserId(cId, uId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cId, uId).Update("num", num).Error
}
