/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:45:49
 */
package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      uint
	Tiitle        string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `goem:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}
