package main

import "fmt"

// 先return 后defer
// defer 是当前函数声明周期全部结束才出栈
func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAnddefer() int {
	defer deferFunc()
	return returnFunc()
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
func main() {
	// 类似于C++的析构函数,如果有多个按照栈来执行
	// defer fmt.Println("main end1")
	// defer fmt.Println("main end2")
	// defer func1()
	// defer func2()
	// defer func3()
	returnAnddefer();
}