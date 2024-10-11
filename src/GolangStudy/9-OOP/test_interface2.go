package main

import "fmt"

// interface 是一种类型

type Foo interface{

}

type Book struct {
	auth string
	name string
}

func myFunc(arg interface{}) {
	fmt.Println(arg)
	// 有类型判断断言机制
	value, ok := arg.(string) // 判断arg是否是string类型的
	if !ok {

	} 
	else {
		// value -> value, type -> %T 
	}
} 

func main() {
	book := Book{"zhagnsan", "CS"}
	myFunc(book)
	myFunc(3.14)


}