package main

import (
	"fmt"
)

// 简单工程模式 实现创建和使用的分离

// 抽象层
type Fruit interface {
	Show()
}

// 实现层
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {

	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

// 工厂模块

type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {

	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit

}

func main() {

	factory := new(Factory)

	//fmt.Println(reflect.TypeOf(factory).String())
	apple := factory.CreateFruit("apple")

	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")

	pear.Show()
}
