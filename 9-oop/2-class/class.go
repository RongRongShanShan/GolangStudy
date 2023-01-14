package main

import "fmt"

/**
go的类方法还是在结构体的基础上进行的
*/

// Book 定义结构体
//	注意！ 类名和其方法首字母大写意味着其他包也可以访问到，而首字母小写(比如Book中定义的两个对象)则视为私有
// 		跨包只能调用其他包中对外开放的方法，(类名、属性名。方法名)首字母大写代表对其他包开放，go语言的封装是针对包来说的，在同一包内大小写没有区别
type Book struct {
	no   int
	name string
}

// Show 声明Book的方法，func后面的()中指明了调用的类型为Book，使用*Book是为了通过引用操作原数据而不是生成的副本(struct不使用引用时会操作生成的副本)
func (b *Book) Show() {
	fmt.Printf("%v\n", *b)
}

func (b *Book) GetNo() int {
	return b.no
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) SetNo(no int) {
	b.no = no
}

func (b *Book) SetName(name string) {
	b.name = name
}

func main() {
	book := Book{no: 1, name: "golang入门"}
	book.Show() // {1 golang入门}
	book.SetName("golang面向对象")
	book.Show() // {1 golang面向对象}
}
