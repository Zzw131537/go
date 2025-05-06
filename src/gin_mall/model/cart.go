/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-16 16:30:49
 */
package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       uint
	MaxNum    uint
	Check     bool
}
