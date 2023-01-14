package main

import "fmt"

/*
变量的定义
*/

// 全局变量声明
var (
	aa int
	bb = 1
)

// :=声明全局变量会报错
//cc := "a"

func main() {
	// 第一种: 指定变量类型但不赋值，使用默认值 (int默认0)
	var a int
	fmt.Println("a = ", a)
	// 第二种: 指定变量类型并赋值
	var b int = 1
	fmt.Println("b = ", b)
	// 第三种: 自动推导类型
	var c = "abc"
	fmt.Printf("c = %s, type = %T\n", c, c)
	// 第四种: 省略var,使用:= (注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误)
	//	此方法不能用于声明全局变量
	d := true
	fmt.Printf("d = %t, type = %T\n", d, d)

	fmt.Println("==== 全局变量 ====")
	fmt.Println("aa = ", aa)
	fmt.Println("bb = ", bb)

	fmt.Println("=== 多变量声明 ===")
	var e, f = 3, "b"
	fmt.Printf("e = %d, f = %s\n", e, f)
}
