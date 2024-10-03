package main

import "fmt"

func main() {	
	// 第一种声明方式
	// 声明myMap是一种map类型 key->string, value->string
	// 底层是hash table
	var myMap map[string]string // = make(map[string]string, 10)
	if myMap == nil {
		fmt.Println("myMap is nil")
	}
	// 先分配空间
	myMap = make(map[string]string, 10)
	myMap["first"] = "C++"
	myMap["second"] = "Go"
	for key, value := range myMap {
		fmt.Println("key = ", key, ",value = ", value)
	}

	// 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "C++"
	myMap2[2] = "Go"
	for key, value := range myMap2 {
		fmt.Println("key = ", key, ", value = ", value)
	}

	// 第三种声明方式
	myMap3 := map[string]string {
		"one" : "C++",
		"two" : "GO",
	}
	for key, value := range myMap3 {
		fmt.Println("key = ", key, ", value = ", value)
	}

}