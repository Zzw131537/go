package main

import "fmt"

// 练习抽象工厂

type AbstracFactory interface {
	CreateXianKa() AbstracXianKa
	CreateNeiChun() AbstracNeiChun
	CreateCpu() AbstracCpu
}
type AbstracXianKa interface {
	ShowXianKa()
}

type AbstracNeiChun interface {
	ShowNeiChun()
}

type AbstracCpu interface {
	ShowCpu()
}

// 实现层
type IntelXianKa struct {
}

func (intelXianKa *IntelXianKa) ShowXianKa() {
	fmt.Println("intelXianKa")
}

type IntelNeiChun struct {
}

func (intelNeiChun *IntelNeiChun) ShowNeiChun() {
	fmt.Println("intelNeiChun")
}

type IntelCpu struct {
}

func (intelCpu *IntelCpu) ShowCpu() {
	fmt.Println("intelCpu")
}

type IntelFactory struct {
}

func (intelFactory *IntelFactory) CreateXianKa() AbstracXianKa {
	return &IntelXianKa{}
}
func (intelFactory *IntelFactory) CreateCpu() AbstracCpu {
	return &IntelCpu{}
}
func (intelFactory *IntelFactory) CreateNeiChun() AbstracNeiChun {
	return &IntelNeiChun{}
}

type KingstonXianKa struct {
}

func (kingstonXianKa *KingstonXianKa) ShowXianKa() {
	fmt.Println("kingstonXianKa")
}

type KingstonCpu struct {
}

func (kingstonCpu *KingstonCpu) ShowCpu() {
	fmt.Println("kingstonCpu")
}

type KingstonNeiChun struct {
}

func (kingstonNeiChun *KingstonNeiChun) ShowNeiChun() {
	fmt.Println("kingstonNeiChun")
}

type KingstonFactory struct {
}

func (kingstonFactory *KingstonFactory) CreateXianKa() AbstracXianKa {
	return &KingstonXianKa{}
}
func (kingstonFactory *KingstonFactory) CreateNeiChun() AbstracNeiChun {
	return &KingstonNeiChun{}
}
func (kingstonFactory *KingstonFactory) CreateCpu() AbstracCpu {
	return &KingstonCpu{}
}

type nvidiaXianka struct{}

func (nvidiaXianka *nvidiaXianka) ShowXianKa() {
	fmt.Println("nvidiaXianKa")
}

type nvidiaNeiChun struct{}

func (nvidiaNeiChun *nvidiaNeiChun) ShowNeiChun() {
	fmt.Println("nvidiaNeiChun")
}

type nvidiaCpu struct{}

func (nvidiaCpu *nvidiaCpu) ShowCpu() {
	fmt.Println("nvidiaCpu")
}

type NvidiaFactory struct{}

func (nvidiaFactory *NvidiaFactory) CreateXianKa() AbstracXianKa {
	return &nvidiaXianka{}
}
func (nvidiaFactory *NvidiaFactory) CreateNeiChun() AbstracNeiChun {
	return &nvidiaNeiChun{}
}
func (nvidiaFactory *NvidiaFactory) CreateCpu() AbstracCpu {
	return &nvidiaCpu{}
}
func main() {

	// 需求一 : Intel 的 3 个硬件

	var intelFactory AbstracFactory

	intelFactory = new(IntelFactory)

	var intelXianKa AbstracXianKa
	intelXianKa = intelFactory.CreateXianKa()
	intelXianKa.ShowXianKa()
	var intelCpu AbstracCpu
	intelCpu = intelFactory.CreateCpu()
	intelCpu.ShowCpu()
	var intelNeiChun AbstracNeiChun
	intelNeiChun = intelFactory.CreateNeiChun()
	intelNeiChun.ShowNeiChun()
	fmt.Println("--------------------------------------------------")

	// 需求二 : intel 的Cpu nvidia 的显卡 Kingston 的内存

	var intelFac AbstracFactory
	var nvidiaFac AbstracFactory
	var kingstonFac AbstracFactory

	intelFac = new(IntelFactory)
	nvidiaFac = new(NvidiaFactory)
	kingstonFac = new(KingstonFactory)

	var xianka AbstracXianKa
	var neicun AbstracNeiChun
	var cpu AbstracCpu

	cpu = intelFac.CreateCpu()
	neicun = kingstonFac.CreateNeiChun()
	xianka = nvidiaFac.CreateXianKa()

	cpu.ShowCpu()
	neicun.ShowNeiChun()
	xianka.ShowXianKa()
	fmt.Println("处理完成！")
}
