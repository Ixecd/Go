package main

import "fmt"

type Human struct {
	name string
	sex string
}

type SuperMan struct {
	Human // 表示SuperMan 类继承了 Human结构体

	
	level int

}


func (this *SuperMan) Eat() {
	fmt.Println("SuperMan's Func -> Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan's Func -> Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
	
}

func (this *Human) Eat() {
	fmt.Println("Human's Func -> Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human's Func -> Walk()...")
}


func main() {
	h := Human{"zhang3", "male"}
	sm := SuperMan{Human{"lisi", "female"}, 1} // 按声明顺序初始化
	var s SuperMan
	s.name = "huawu"
	s.sex = "female"
	s.level = 1
	s.Print()
	h.Eat()
	h.Walk()
	sm.Eat()
	sm.Fly()
}