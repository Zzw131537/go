/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 19:38:51
 */
package model

import (
	"mall/cache"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Tiitle        string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViwKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViwKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
