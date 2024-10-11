package main

import "fmt"

func main() {
	myArray := []int{}
	// 长度为3,容量为5
	myArray = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(myArray),cap(myArray), myArray)

	// 向myArray切片追加一个元素
	myArray = append(myArray, 10)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(myArray),cap(myArray), myArray)

	myArray = append(myArray, 100)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(myArray),cap(myArray), myArray)

	// 2倍扩增
	myArray = append(myArray, 1000)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(myArray),cap(myArray), myArray)

}