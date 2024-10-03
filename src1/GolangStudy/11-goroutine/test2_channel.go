package main

import "fmt"
import "time"

func main() {
	// 带有缓冲
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))
	// 天生的并发
	go func() {
		defer fmt.Println("child goroutine over...")

		// 由于chann中cap为3 所以当i = 3 时会阻塞
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("child goroutinue running,cur i = ", i, ",len(c) = ", len(c), ", cap(c) = ", cap(c))
		} 
	}()

	//time.Sleep(1 * time.Second)

	for i := 0; i < 4; i++ {
		// 这里 <-c 之后chann有空闲位置,上面的阻塞会解除
		num := <-c
		// time.Sleep(1 * time.Second)
		fmt.Println("num = ", num)
	}

	fmt.Println("main over...")

}