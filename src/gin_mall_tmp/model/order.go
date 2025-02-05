/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:53:16
 */
package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BissId    uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderNum  int64
	Type      int // 1 未支付 2 已支付
	Money     float64
}
