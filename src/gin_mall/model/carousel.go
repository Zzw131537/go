/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:06:26
 */
package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
