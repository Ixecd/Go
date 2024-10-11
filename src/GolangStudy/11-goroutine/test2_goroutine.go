package main

import (
	"fmt"
	"time"
)

func main() {
	// go 承载一个 匿名函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}() // 定义且调用

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
	// 使用runtime.Goexit()来退出协程

}