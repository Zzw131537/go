/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 19:57:14
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(da *gorm.DB) *AddressDao {
	return &AddressDao{da}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByaId(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByuId(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id = ?", uId).Find(&addresses).Error
	return
}

func (dao *AddressDao) UpdateAddressByUserId(address *model.Address, uId, aId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id = ? AND user_id =?", aId, uId).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddress(uId, aId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id=? AND user_id = ? ", aId, uId).Delete(&model.Address{}).Error
}
