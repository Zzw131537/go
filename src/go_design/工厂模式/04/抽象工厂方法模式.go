package main

import "fmt"

// 降低工厂的数量
// 等级结构需要稳定
// 针对产品族进行开发符合开闭，针对产品等级结构进行开发会影响开闭原则

// 抽象层
type AbstracApple interface {
	ShowApple()
}

type AbstracBananc interface {
	ShowBanana()
}

type AbstracPear interface {
	ShowPear()
}

// 抽象工程
type AbstraFactory interface {
	CreateApple() AbstracApple
	CreateBanana() AbstracBananc
	CreatePear() AbstracPear
}

// 实现层

// 中国产品
type ChinaApple struct {
}

func (chinaApple *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {
}

func (chinaBanana *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct {
}

func (chinaPear *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct {
}

func (chainFactory *ChinaFactory) CreateApple() AbstracApple {
	return &ChinaApple{}
}

func (chainFactory *ChinaFactory) CreateBanana() AbstracBananc {
	return &ChinaBanana{}
}

func (chainFactory *ChinaFactory) CreatePear() AbstracPear {
	return &ChinaPear{}
}

// 日本产品族
type JapanApple struct {
}

func (japanApple *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct {
}

func (japanBanana *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanPear struct {
}

func (japanPear *JapanPear) ShowPear() {
	fmt.Println("日本梨")
}

type JapanFactory struct {
}

func (japanFactory *JapanFactory) CreateApple() AbstracApple {
	return &JapanApple{}
}
func (japanFactory *JapanFactory) CreateBanana() AbstracBananc {
	return &JapanBanana{}
}
func (japanFactory *JapanFactory) CreatePear() AbstracPear {
	return &JapanPear{}
}

// 逻辑层
func main() {

	// 需求一 ： 需要中国的水果

	var cFac AbstraFactory
	cFac = new(ChinaFactory)

	var cApple AbstracApple

	cApple = cFac.CreateApple()
	cApple.ShowApple()

	var cBanana AbstracBananc
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()

	var cPear AbstracPear
	cPear = cFac.CreatePear()
	cPear.ShowPear()
	fmt.Println("--------------------------------------------------")
	// 需求二 ： 需要日本水果

	var jFac AbstraFactory
	jFac = new(JapanFactory)

	jFac = new(JapanFactory)
	var jApple AbstracApple
	jApple = jFac.CreateApple()
	jApple.ShowApple()
	var jBanana AbstracBananc
	jBanana = jFac.CreateBanana()
	jBanana.ShowBanana()
	var jPear AbstracPear
	jPear = jFac.CreatePear()
	jPear.ShowPear()
	fmt.Println("--------------------------------------------------")
}
