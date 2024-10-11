package main

import(
	"fmt"
	"time"
)

// 全局变量
var gA int = 100
var gB = "hello gB"
// := 只能用在函数体内
// gC := 100


func main() {
	fmt.Println("hello Go!")
	time.Sleep(1 * time.Second)

	// 声明变量
	var a int;
	// Println中ln表示换行输出, f表示格式化输出
	fmt.Println("a = ", a)

	var b int = 100;
	// error -> var b int 100 赋初值必须加=

	fmt.Println("b = ", b)

	// 在初始化的时候可以省去数据类型的声明

	var c = "1234";
	fmt.Println("c = ", c)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	// 可以打印类型的数据类型
	fmt.Printf("type of a = %T\n", a)
	fmt.Printf("type of b = %T\n", b)
	fmt.Printf("type of c = %T\n", c)


	// 可以省区var关键字,直接自动匹配
	e := 100
	fmt.Printf("e = %d, type of e = %T\n", e, e)

	g := 3.14
	fmt.Println("g = ", g);
	fmt.Printf("type of g = %T\n", g)


	//fmt.Println("gC = ", gC)
	//fmt.Printf("type of gC = %T\n", gC)

	fmt.Println("gB = ", gB)
	fmt.Printf("type of gB = %T\n", gB)

	// 声明多个变量
	var aa, bbb int = 100, 200
	fmt.Println("aa = ", aa, " bbb = ", bbb);

	// 多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj);

}
