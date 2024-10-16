package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")	
} 

func (this *Book) WriteBook() {
	fmt.Println("Write a book")	
} 

func main() {

	var a string = "lars" // -> pair<static type::string, value:"lars">

	fmt.Println(a)
	// pair<type::string, value:"lars">
	var allType interface{}
	allType = a

	fmt.Printf("type of alltype: %T\n", allType)

	str, _ := allType.(string)
	fmt.Println(str)

	b := &Book{}

	// pair<type: , value: >
	var r Reader;
	// r: pair<type::Book, value:Book{}的地址>

	// b -> 派生类,  r -> 基类
	r = b 

	r.ReadBook()
	b.WriteBook()
	// pair<type: , value: >
	var w Writer;
	// w:
	w = r.(Writer) // w 和 r具体的type是一样的

	fmt.Printf("type of w and r: %T, %T\n", w, r)

	w.WriteBook()

}