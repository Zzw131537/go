/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 17:20:26
 */
package dao

import (
	"context"
	"mall/model"

	"gorm.io/gorm"
)

type CarouseleDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouseleDao {
	return &CarouseleDao{NewDBClient(ctx)}
}

func NewCarouseDaoByDB(da *gorm.DB) *CarouseleDao {
	return &CarouseleDao{da}
}

func (dao *CarouseleDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
