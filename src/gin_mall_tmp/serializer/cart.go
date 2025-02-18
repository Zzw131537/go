/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-16 16:49:17
 */
package serializer

import (
	"context"
	"mall/dao"
	"mall/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	name          string `json:"name"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		CreatedAt:     cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		Check:         cart.Check,
		name:          product.Name,
		DiscountPrice: product.DiscountPrice,
		ImgPath:       product.ImgPath,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.BossId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)

	}
	return
}
