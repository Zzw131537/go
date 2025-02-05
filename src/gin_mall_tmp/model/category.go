/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:19:05
 */
package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string
}
