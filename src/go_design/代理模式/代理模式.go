package main

import "fmt"

type Goods struct {
	Kind string // 种类
	Fact bool   // 真伪
}

// 抽象层
type Shopping interface {
	Buy(goods *Goods)
}

// 实现层
type ChinaShopping struct {
}

func (c *ChinaShopping) Buy(goods *Goods) {
	fmt.Println("中国购买"+goods.Kind, "商品！")
}

type JapanShopping struct {
}

func (j *JapanShopping) Buy(goods *Goods) {
	fmt.Println("日本购买"+goods.Kind, "商品！")
}

// 海外代理
type OverSeaProxy struct {
	shopping Shopping
}

func main() {

}
