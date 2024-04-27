package main

import "fmt"


func foo(a int, b string) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c;
}

// go 允许多个返回值(匿名)
func foo2(a int, b string) (int, string) {
	return a, b
}
// 有形参名(默认为0)
func foo3(a string, b int) (r1 int,r2 int) {
	return 100, 200
}

func main() {

	d := foo(1, "22");
	fmt.Println("d = ", d)

	var aa, bb = foo2(1, "22");
	fmt.Println("aa = ", aa, "bb = ", bb)

	var cc, dd = foo3("22", 1);
	fmt.Println("cc = ", cc, "dd = ", dd)

	var c int
	fmt.Println("c = ", c)
	
}