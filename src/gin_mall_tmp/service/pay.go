/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 21:51:54
 */
package service

import (
	"context"
	"errors"
	"fmt"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"strconv"
)

type OrderPayService struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"`
}

func (service *OrderPayService) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	_, err := orderDao.GetOrderById(service.OrderId, uId)
	if err != nil {
		util.LogrusObj.Infoln("err ", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	money := service.Money
	num := service.Num
	money = money * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Infoln("err ", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱进行解密
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}

	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	var boss *model.User
	boss, err = userDao.GetUserById(uint(service.BossId))
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}

	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)

	err = userDao.UpdateUserById(boss.ID, boss)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("出现错误").Error(),
		}
	}

	// 对应的商品数目减一
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)
	product.Num -= num
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 订单shancg
	err = orderDao.DeleteOrder(uId, service.OrderId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 自己的商品商品 + 1
	productUser := &model.Product{
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Tiitle:        product.Tiitle,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        false,
		Num:           1,
		BossId:        uId,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}
	err = productDao.CreateProduct(productUser)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
