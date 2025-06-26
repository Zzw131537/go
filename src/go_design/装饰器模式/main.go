package main

import "fmt"

// 装饰器 模式 比较灵活
// 抽象层
type Phone interface {
	Show()
}

// 抽象的装饰器
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// 实现层

type HuaWei struct{}

func (h *HuaWei) Show() {
	fmt.Println("华为手机")
}

type XiaoMi struct{}

func (x *XiaoMi) Show() {
	fmt.Println("小米手机")
}

// 具体的装饰器
type MoDecorator struct {
	Decorator
}

func (m *MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("贴膜")
}

func NewMoDecorator(phone Phone) *MoDecorator {
	return &MoDecorator{Decorator{phone: phone}}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) Show() {
	k.phone.Show()
	fmt.Println("加壳")
}
func NewKeDecorator(phone Phone) *KeDecorator {
	return &KeDecorator{Decorator{phone: phone}}
}

// 业务逻辑层
func main() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()
	fmt.Println("--------------------------------------------------")
	var f Phone
	f = NewMoDecorator(huawei)
	f.Show()
	fmt.Println("--------------------------------------------------")
	var kemohuawei Phone
	kemohuawei = NewKeDecorator(NewMoDecorator(f))
	kemohuawei.Show()
}
