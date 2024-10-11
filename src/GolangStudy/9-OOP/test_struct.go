package main

import "fmt"

// 声明一种新的数据类型,int的一个别名
type myint int

type Book struct {
	title string
	auto string
}

// 如果说类的属性首字母为大写,表示公开,否则只能内部访问
// 方法同理,大写表示对外开放
type Hero struct {
	Name string
	Ad int
	Level int// 延续上面的类型
}

// // -> this Hero 表示当前方法绑定到Hero结构体
// func (this Hero) GetName() {
// 	fmt.Println("name = ", this.Name)
// }
// // 当前this是值拷贝
// func (this Hero) SetName() {
// 	this.Name = "Hero"
// }

// -> this Hero 表示当前方法绑定到Hero结构体
// func(s StrcutName) funcName() {} --> 其中(s StructName) 表示这个方法属于StructName,可以通过s来访问实例化后的struct内容
func (this *Hero) GetName() {
	fmt.Println("name = ", this.Name)
}
// 当前this是值拷贝
func (this *Hero) SetName() {
	this.Name = "Hero"
}


// 值传递
func changeBook(book Book) {
	book.title = "nihao"
}
// 指针传递
// func changeBook(book *Book) {
// 	*book.title = "noo"
// }

func main() {

	var book1 Book
	book1.title = "Computer Science"
	book1.auto = "..."

	fmt.Printf("%v\n", book1)
	 
	// Go中的struct中的this当作struct来处理
	// 就算传入的是*this 也当struct来处理
	hero := Hero{Name : "123", Ad : 123, Level : 1}
	hero.SetName()
	hero.GetName()
}