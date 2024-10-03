package main

// 函数名首字母大写表示对外开放,否则只能在包内使用
// 导包就得用,如果不想用就要起别名
import (
	_ "GolangStudy/5-import-init/lib1" // 匿名
	"GolangStudy/5-import-init/lib2" // 直接导入到当前包
)
/*
	import _ "fmt" 表示给fmt起一个别名,匿名,无法使用当前包的方法,但是会执行当前的包内部的init()方法
	import aa "fmt" 给fmt起一个别名aa,aa.Println()来使用
	import . "fmt" 将fmt包中的全部方法,导入到当前包内
*/
func main() {
	//lib1.Lib1Test();
	lib2.Lib2Test();
}