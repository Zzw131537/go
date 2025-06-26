package main

// 将两个不同的种类，进行连接协作

// 适配的目标
type VS interface{}

type Phone struct {
	v VS
}

func NewPhone(v VS) *Phone {
	return &Phone{v}
}

func main() {

}
