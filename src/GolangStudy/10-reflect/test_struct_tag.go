package main

import "fmt"
import "reflect"

// 加一个tag -> 标签 key - value 形式
type resume struct {
	Name string	 `info:"name" doc:"myName"` 
	Sex string	`info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() // 当前结构体的所有元素

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagstring, "doc: ", tagdoc)
	}
}

func main() {
	var re resume

	findTag(&re)

}