/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 21:23:07
 */
package service

import (
	"context"
	"mall/dao"
	"mall/serializer"
	"strconv"
)

type ProductImgService struct {
}

func (service *ProductImgService) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
