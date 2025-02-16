/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-16 15:09:14
 */
package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Addtrss string `gorm:"type:varchar(50) not null"`
}
