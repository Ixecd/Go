package main

import "fmt"
import "reflect" // reflect package

type User struct {
	Id int
	Name string
	Age int
}


// 这两个不能同时存在
func (this User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

// func (this *User) Call() {
// 	fmt.Println("*user is called...")
// 	fmt.Printf("%v\n", this)
// }

func reflectNum(arg interface{}) {
	fmt.Println("type = ", reflect.TypeOf(arg))
	fmt.Println("value = ", reflect.ValueOf(arg))	
}

func DoDiledAndMethod(input interface{}) {
	// 获取input的type main.User 而非main.*User
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType = ", inputType)
	// 同时获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue = ", inputValue)
	// 分别通过type获取里面的字段
	// 1.获取interface的type类型,根据type得到field,进行遍历
	// 2.得到每个field,数据类型
	// 3.通过filed有一个Interface()方法得到对应的value
	for i:=0; i < inputType.NumField(); i++ {
		field := inputType.Field(i) // 取出第i个field field.Name() field.Type()
		value :=inputValue.Field(i).Interface()

		//fmt.Println("i = ", i , ", field = ", field, ", value = ", value)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	var num float64 = 3.14159

	reflectNum(num)

	user := User{1, "lars", 18}

	DoDiledAndMethod(user)

}