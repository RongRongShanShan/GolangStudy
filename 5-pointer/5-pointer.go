package main

/**
指针
*/

// 通过指针交换形参的值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	a, b := 1, 2
	//引用传递 (传递地址)
	swap(&a, &b)
	println(a, b)

	//指针类型
	var p *int = &a
	println(p) // 打印结果是a的地址

	//二级指针
	var pp **int = &p
	println(pp) // 打印结果是p的地址
}
