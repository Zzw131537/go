/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:02:03
 */
package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Addtrss string `gorm:"type:varchar(50) not null"`
}
