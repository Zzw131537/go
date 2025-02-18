/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:35:03
 */
package service

import (
	"context"
	"fmt"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  int     `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var order *model.Order
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, // 默认未支付

	}

	// 检验地址存不存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByaId(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	order.AddressId = address.ID

	// 生成订单号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = int64(orderNum)

	err = orderDao.CreateOrder(order)

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

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uint(orderId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByaId(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(order.ProductId))
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
		Data:   serializer.BuildOrder(order, product, address),
	}

}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	if service.Type != 1 && service.Type != 2 {
		service.Type = 0
	}

	orders, err := orderDao.ListOrderByCondition(uint(service.Type), uId, service.BasePage)

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
		Data:   serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(len(orders))),
	}
}

func (service *OrderService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	orderId, _ := strconv.Atoi(cId)

	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	err := orderDao.DeleteOrder(uId, uint(orderId))
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
