package main 
import "fmt"

func main() {
	fmt.Println("Hello World!");
}

/*
	针对上面的代码对调度器里面的结构做一个分析

	1. runtime创建最初的thread和processor,并将两者关联
	2. 调度器初始化:初始化thread/栈/垃圾回收,以及创建和初始化由GOMAXPROCS个Processor构成的processor列表
	3. 上面的代码中main函数是`main.main`, runtime中也有一个main函数,runtime.main,代码经过编译后,runtime.main会调用main.main,程序启动时会为runtime.main创建gorutine,之后把其加入到Processor本地队列中
	4. 启动thread,这时候第一个线程已经绑定了processor,会从processor的本地队列中获取goroutine
	5. goroutine拥有栈,thread根据goroutine中的栈信息和调度信息设置运行环境
	6. thread运行goroutine(runtime.mian中创建的)
	7. goroutine执行完毕,返回到thread,由thread通过processor从队列中获取goroutine执行,如果没有任务,就会将本线程挂起,有任务就会唤醒
	8. 一直循环下去,直到程序执行完毕,也就是main.main退出,runtime.main执行defer和panic处理,或者调用runtime.exit退出程序

*/

