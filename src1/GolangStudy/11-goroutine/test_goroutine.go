package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine : i = ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// 创建一个go协程, 去执行newTask()
	go newTask()

	i := 0
	for {
		i++;
		fmt.Println("main Func : i = ", i)
		time.Sleep(1 * time.Second)
	}
}