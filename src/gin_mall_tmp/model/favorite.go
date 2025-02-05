/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:31:33
 */
package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `grom:"not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `grom:"not null"`
	Boss      User    `gorm:"ForeignKey:"BossId"`
	BossId    uint    `gorm:"not null"`
}
