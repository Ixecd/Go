package main

import "fmt"

func printArray(myArray []int) {
	// slice是引用传递
	// _ 表示匿名,无法访问
	for _, value := range myArray {
		fmt.Println(value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{11,22,33,44}
	fmt.Printf("myArray's type = %T\n", myArray)
	printArray(myArray)
	printArray(myArray)
}