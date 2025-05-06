/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:03:32
 */
package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string
}
