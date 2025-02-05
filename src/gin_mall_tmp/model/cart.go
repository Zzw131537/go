/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:08:44
 */
package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"`
	Check     bool `gorm:"not null"`
}
