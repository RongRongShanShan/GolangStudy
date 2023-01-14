package main

/**
init函数与import
	init函数可在package main中，也可在其他package中，并且可在同一个package中出现多次
*/

/**
导包 (会调用init方法)
	import 别名 包名  - 可以使用别名使用包中的函数
	import _ 包名    - 导入的包没有使用时使用这种格式
	import . 包名    - 使用包中函数时可以不写包名直接使用 (不建议，不同包同名函数会有歧义)
*/
import (
	"GolangStudy/4-init/lib1"
	. "GolangStudy/4-init/lib2"
)

func init() {
	println("4-init.go 的 init方法被调用")
}

func main() {
	println("main 方法 被调用")
	// 其他包下的小写开头的方法没法调用，报错为未导出
	lib1.Fun()
	//lib2.Fun()
	Fun()
}

/**
结果如下：

lib1 - init 被调用
lib2 - init 被调用
4-init.go 的 init方法被调用
main 方法 被调用
lib1 中的方法 被调用
lib2 中的方法 被调用
*/
