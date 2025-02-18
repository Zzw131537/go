/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 18:25:46
 */
package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var cart *model.Cart
	// 判断有没有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
	}
	err = cartDao.CreateCart(cart)

	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	carts, err := CartDao.ListCartByuId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

// 只跟新数量
func (service *CartService) Update(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)

	code := e.Success
	cartDao := dao.NewCartDao(ctx)

	err := cartDao.UpdateCartNumBuUserId(uint(cartId), uId, service.Num)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)

	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCart(uId, uint(cartId))
	if err != nil {
		util.LogrusObj.Infoln("err ", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
