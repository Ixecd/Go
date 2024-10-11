package main

import "fmt"


// 关于interface --> 通用类型 --> interface{} -> 空接口、int string float64 struct
// 可以使用interface{}类型引用 任意数据类型

// 多态 -> 有父类(接口) interface
// 继承 -> 有子类(实现所有父类的接口方法)
// 多态 -> 基类指针指向派生类对象


// interface 类型本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	// 关于interface,不需要显示的继承,只需要在结构体中实现所有interface中的函数,就相当于继承了interface->类似于C++中的virtual fun() = 0
	// 注意:全部
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat's Sleep()...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 不属于任何一个类
func showAnimal(animal AnimalIF) {
	animal.Sleep() // -> 实现多态,类似于C++中的基类指针指向派生类对象,根据虚函数表调用相关函数
}

func main() {

	// 接口的数据类型,父类指针
	var animal AnimalIF
	// 初始化不能使用(),只能使用{}	
	animal = &Cat{"Green"}

	animal.Sleep()
	animal.GetColor()
	fmt.Println(animal.GetType())

}