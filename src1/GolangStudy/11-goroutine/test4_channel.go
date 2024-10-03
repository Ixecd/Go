package main

import "fmt"


func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c<-i
			fmt.Println("send in channel i = ", i)
		}
		close(c)
	}()
	// range阻塞的等待c的结果
	for data := range c {
		fmt.Println("data = ", data);
	}

	fmt.Println("Main fun end...")


}