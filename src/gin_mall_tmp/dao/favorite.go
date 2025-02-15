/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-15 17:26:24
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(da *gorm.DB) *FavoriteDao {
	return &FavoriteDao{da}
}

func (dao *FavoriteDao) ListFavorite(uId uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uId).Find(&resp).Error
	return

}

func (dao *FavoriteDao) FavoriteExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND id = ?", uId, fId).Delete(&model.Favorite{}).Error
}
