package main

import "fmt"

// 对于工厂模式不必更改原代码
// 符合开闭原则，但增加了工作量
// 抽象层

type Fruit interface {
	Show()
}

// 工厂类
type AbstracFactory interface {
	CreateFruit() Fruit
}

// 继承工厂层
type AppleFactory struct {
	AbstracFactory
}

type BananaFactory struct {
	AbstracFactory
}

type PearFactory struct {
	AbstracFactory
}

func (appleFactory *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

func (bananaFactory *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

func (pearFactory *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 基础模块层
type Apple struct {
	Fruit
}

type Banana struct {
	Fruit
}

type Pear struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}
func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}
func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

func main() {

	// 需求一 需要一个具体对象

	var factory AppleFactory
	var apple = factory.CreateFruit()
	apple.Show()

	var banFac BananaFactory

	var banana = banFac.CreateFruit()
	banana.Show()

	var pearFac PearFactory

	var pear = pearFac.CreateFruit()
	pear.Show()

}
