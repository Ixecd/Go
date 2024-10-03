package main

import "fmt"


func main() {
	myArray := []int{11,22,33,44}
	
	// %v -> 打印全部信息
	fmt.Printf("len = %d, slice = %v\n",len(myArray), myArray)
	
	var slice1 []int
	// 通过make来重新构建大小
	slice1 = make([]int, 3)
	slice1[0] = 100

	// 也可以同时进行
	var slice []int = make([]int, 5)
	fmt.Printf("slice's type = %T\n", slice)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	for _, value := range slice1 {
		fmt.Println("value = ", value)
	}
	

	// ---------- 总结 ----------
	// 1. slice1 := []int{1, 2, 3, 4}
	// 2. var slice2 []int 
	// 	  slice2 = make([]int, 5)
	// 3. var slice3 []int = make([]int, 5)
	// 4. slice4 := make([]int, 6)

	// 如何判断一个slice是否为空
	if slice == nil {
		fmt.Println("slice 空切片")
	}

}