package main

// 注意使用GOPATH来找包的这种方式已经被抛弃了,所以后面都是使用go.mod来管理依赖包

// 函数名首字母大写表示对外开放,否则只能在包内使用
// 导包就得用,如果不想用就要起别名
import (
	// 注意,如果直接这样用的话,必须关闭模块化管理,否则会报错
	// 设置 GO111MODULE=off
	// go env -w GO111MODULE=off
	// 设置 GOPATH
	// go env -w GOPATH="/home/qc/Go"
	// 这里搜索的过程就是先从标准库中找,如果没有找到就从GOPATH中找,GOPATH是之前已经设置好的,从GOPATH的src下面开始找,所以这里的"GolangStudy..."不需要写全路径
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