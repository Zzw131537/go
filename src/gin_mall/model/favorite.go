/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:31:33
 */
package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"foreignKey:"BossId"`
	BossId    uint    `gorm:"not null"`
}
