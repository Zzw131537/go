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
	Name          string // 商品名称
	CategoryId    uint   // 所属种类
	Tiitle        string // 标题
	Info          string //信息
	ImgPath       string // 图片
	Price         string // 价格
	DiscountPrice string // 打折后的价格
	OnSale        bool   `gorm:"default:false"` // 状态
	Num           int    // 库存
	BossId        uint   // 商家id
	BossName      string // 商家姓名
	BossAvatar    string // 商家头像
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
