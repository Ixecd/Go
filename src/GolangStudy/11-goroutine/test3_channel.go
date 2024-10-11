package main

import "fmt"

func main() {

	c := make(chan int)
	// 不能对同一个channel多次close()否则panic
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 关闭一个channel
		close(c)
	}()
	
	for {
		// ok如果为true表示channel没有关闭
		// 判断ok -> ok指的是管道有没有关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			fmt.Println("对端已关闭")
			break;
		}

	}

}