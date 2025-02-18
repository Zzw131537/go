/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-18 20:18:48
 */
package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint
	AddressId uint
	Num       int
	OrderNum  int64
	Type      int // 1 未支付 2 已支付
	Money     float64
}
