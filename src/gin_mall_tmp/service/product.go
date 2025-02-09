/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 18:16:31
 */
package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"mime/multipart"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)

	// 以第一·张作为封面图
	tmp, _ := files[0].Open()
	path, err := UpLoadproductToLocalStatic(tmp, uId, service.Name)
}
