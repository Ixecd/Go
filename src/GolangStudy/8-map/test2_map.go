package main

import "fmt"

// 这里是引用传递
func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func main() {

	// 这里 var name = 
	// 或者 name := 
	myMap := make(map[string]string)

	myMap["first"] = "C++"
	myMap["second"] = "Go"

	for key, value := range myMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	// 修改
	myMap["second"] = "Java"
	for key, value := range myMap {
		fmt.Println("key = ", key, "value = ", value)
	}

	// 删除
	delete(myMap, "second")

	for key, value := range myMap {
		fmt.Println("key = ", key, "value = ", value)
	}

}