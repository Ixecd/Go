package main

import "fmt"

// const 可以修饰枚举类型
const (
	// iota 只能够配合const() 一起使用
	// iota 每行都会累加一,默认第一行为0,仅是iota每行会加一,按表达式的结果来赋值(按行来)
	BEIJING = 10 * iota
	SHANGHAI 	// -> 10 * iota = 10
	SHENZHEN	// -> 10 * iota = 20
)

const (
	a, b = iota + 1, iota + 2 	// -> iota = 0, a = 1, b = 2
	c, d						// -> iota = 1, c = 1 + 1 = 2, d = 1 + 2 = 3
	e, f						// -> iota = 2, e = 2 + 1 = 3, f = 2 + 2 = 4

	g, h = iota * 2, iota * 3 	// -> iota = 3, g = 6, h = 9
	i, k						// -> iota = 4, i = 8, k = 12
)

func main () {
	const length int = 100
	fmt.Println("length = ", length)
	fmt.Printf("type of length = %T\n", length)
	fmt.Println("B = ", BEIJING, " SH = ", SHANGHAI, " SZ = ", SHENZHEN)

}