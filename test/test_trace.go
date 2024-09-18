package main

import (
	"os"
	"fmt"
	"runtime/trace"
)


// go tool tarce
func main() {
	// 创建trace文件
	f, err := os.Create("trace.out");
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	fmt.Println("Hello World!");

}
