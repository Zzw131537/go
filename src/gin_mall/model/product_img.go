/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:48:36
 */
package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}
