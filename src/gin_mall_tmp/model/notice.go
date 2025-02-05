/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:49:46
 */
package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
