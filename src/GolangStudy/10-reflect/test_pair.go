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

	
	var a string = "lars" // -> pair<statictype::string, value:"lars">

	fmt.Println(a)
	// pair<type::string, value:"lars">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

	b := &Book{}

	// pair<type: , value: >
	var r Reader;
	// r: pair<type::Book, value:Book{}的地址>
	r = b 

	r.ReadBook()
	b.WriteBook()
	// pair<type: , value: >
	var w Writer;
	// w:
	w = r.(Writer) // w 和 r具体的type是一样的

	w.WriteBook()

}