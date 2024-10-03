package main

import "fmt"

// 长度为4的和长度为10的不是同一种数组
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
		
}

func main() {
	// 固定长度数组
	var myArray [10]int

	for i := 0; i < 10; i++ {
		fmt.Println(myArray[i])
	}

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{2,5,6,8}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	fmt.Printf("myArray's type = %T\n", myArray)
	fmt.Printf("myArray3's type = %T\n", myArray3)

	printArray(myArray3)
	
}