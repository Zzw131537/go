/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 17:24:30
 */
package service

import (
	"context"
	"mall/dao"
	"mall/pkg/e"
	"mall/serializer"
)

type CarouselService struct {
}

// 轮播图
func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)

	code := e.Success

	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
