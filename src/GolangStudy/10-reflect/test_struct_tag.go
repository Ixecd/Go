package main

import "fmt"
import "reflect"

// 加一个tag -> 标签 key - value 形式
type resume struct {
	Name string	 `info:"name" doc:"myName"` 
	Sex string	`info:"sex"`
}

func findTag(str interface{}) {
	// 如果这里传入的是&re表示re的地址,此时TypeOf(str)表示指针类型,其中的Elem()才表示具体的结构体类型.
	t := reflect.TypeOf(str).Elem() 

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagstring, "doc: ", tagdoc)
	}
}

func main() {
	var re resume

	findTag(re) // -> 传值,就不需要 t := reflect.TypeOf(str).Elem() 了
	findTag(&re) // -> 传指针

}