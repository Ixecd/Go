package main

import "fmt"

// chan 有同步机制

func main() {
	// chan 的类型为 int
	// 无缓存
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束...")

		fmt.Println("goroutine 正在运行...")

		c <- 666 // 将666写到c中
	}()

	num := <- c // 从c中接受数据,并且存储再num中 
	
	fmt.Println("num = ", num)
	fmt.Println("main goroutine over...")
}