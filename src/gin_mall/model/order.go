/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:18:48
 */
package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserId    uint    `gorm:"not null"` // 下订单的用户id
	ProductId uint    `gorm:"not null"` // 商品id
	BossId    uint    //商家id
	AddressId uint    // 用户要填的地址id
	Num       int     // 数量
	OrderNum  int64   // 订单号
	Type      int     // 1 未支付 2 已支付
	Money     float64 // 所要付的钱
}
