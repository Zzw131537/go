package main

import (
	"fmt"
	"sync"
)

type singektion struct {
}

var instance *singektion

// 全部私有化，外部怎么访问
// 对外提供方法获取对象

// 只有读权限 , 懒汉模式

var lock sync.Mutex

var initalized uint32

var once sync.Once

func GetInstance() *singektion {

	// 原子操作
	// if atomic.LoadUint32(&initalized) == 1 {
	// 	return instance
	// }

	// // 懒汉式 ，线程安全
	// lock.Lock()

	// defer lock.Unlock()
	// if instance == nil {
	// 	// 原子加 1
	// 	instance = &singektion{}
	// 	atomic.StoreUint32(&initalized, 1)
	// }

	once.Do(func() {
		instance = &singektion{}
	})

	return instance
}

func (s *singektion) SomeThing() {
	fmt.Println("单例模式")
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("单例模式成功")
	} else {
		fmt.Println("单例模式失败")
	}
}
