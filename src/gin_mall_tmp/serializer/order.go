/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:27:33
 */
package serializer

import (
	"context"
	"fmt"
	"mall/conf"
	"mall/dao"
	"mall/model"
)

type Order struct {
	Id            uint    `json:"id"`
	OrderNum      uint    `json:"order_num"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
	UserId        uint    `json:"user_id"`
	ProductId     uint    `json:"product_id"`
	BossId        uint    `json:"boss_id"`
	Num           int     `json:"num"`
	addressName   string  `json:"address_name"`
	AddressPhone  string  `json:"address_phone"`
	Address       string  `json:"address"`
	Type          int     `json:"type"`
	ProductName   string  `json:"product_name"`
	ImgPath       string  `json:"img_path"`
	DiscountPrice string  `json:"discount_price"`
	Money         float64 `json:"money"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		Id:          order.ID,
		OrderNum:    uint(order.Num),
		CreatedAt:   order.CreatedAt.Unix(),
		UpdatedAt:   order.UpdatedAt.Unix(),
		UserId:      order.UserId,
		ProductId:   order.ProductId,
		BossId:      order.BossId,
		Num:         order.Num,
		addressName: address.Name,
		Address:     address.Addtrss,
		Type:        order.Type,
		ProductName: product.Name,
		ImgPath:     conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Money:       order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	AddressDao := dao.NewAddressDao(ctx)
	fmt.Println("长度为:", len(items))
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := AddressDao.GetAddressByaId(item.AddressId)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return
}
