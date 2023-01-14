package main

import "fmt"

/**
函数
*/

// 函数简单示例
func foo1(a int, b string) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return a + len(b)
}

// 多返回值 (匿名)
func foo2(a int, b string) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return a, len(b)
}

// 多返回值 (返回值有形参名)
func foo3(a int, b string) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 给返回值形参赋值 (注意 r1和r2属于foo3的形参，初始值为默认值0)
	r1, r2 = a, len(b)
	// 直接return，也可以不赋值将结果return(就是不赋值直接 return a, len(b) )
	return
}

// 多返回值 (返回值形参类型相同时可以合并)
func foo4(a int, b string) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1, r2 = a, len(b)
	return
}

// 参数类型相同也可以合并
func foo5(a, b int) (r1, r2 int) {
	return //直接返回就会返回r1,r2的默认值0
}

func main() {
	c := foo1(1, "abc")
	fmt.Println("c = ", c) // c =  4

	fmt.Println("=== 多返回值(匿名) ===")
	d, e := foo2(1, "abc")
	fmt.Println("d = ", d, "  e = ", e) // d =  1   e =  3

	fmt.Println("=== 多返回值(返回值有形参名) ===")
	d, e = foo3(1, "abc")
	fmt.Println("d = ", d, "  e = ", e) // d =  1   e =  3

	fmt.Println("=== 多返回值(形参类型相同时可以合并) ===")
	d, e = foo4(1, "abc")
	fmt.Println("d = ", d, "  e = ", e) // d =  1   e =  3
}
