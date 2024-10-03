package main

import "fmt"

func func1(a int, b string) {
	fmt.Println("a = ", a, ", b = ", b)
}

// func func1(a string, b int) {
// 	fmt.Println("a = ", a, ", b = ", b)
// }

func main() {

	var a int = 100
	var b string = "qc"
	func1(a, b)
	//func1(b, a)
	
}