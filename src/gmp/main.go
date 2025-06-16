package main

/*
				调度器的设计策略
				复用线程: workStealing

			    handoff

				利用并行

			go func() -> 创建G -> 入本地队列 -> 入全局队列


		go 的启动周期
		M0: 启动程序后编号为0的主线程
		G0:每启动一个M,都会创建第一个goroutine ,中间转换的桥梁

	可视化的GMP 调试

通过 go tool trace name
*/

func main() {

	//// 1.创建trace文件
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer f.Close()
	//
	//// 启动
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("hello world")
	//// 停止trace
	//trace.Stop()

	// 场景1

}
