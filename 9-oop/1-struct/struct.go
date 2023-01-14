package main

import "fmt"

/**
struct结构体
*/

// type关键字, 声明一种数据类型，这里是给int起了一个别名
type myint int

// Book 定义一个结构体Book
type Book struct {
	no   int
	name string
}

func main() {
	var book1 Book
	book1.no = 1
	book1.name = "《java并发编程的艺术》"
	fmt.Printf("%v\n", book1) // {1 《java并发编程的艺术》}

	//struct作为函数参数传递的是副本，在函数中不会修改原值，需要在函数中修改同样使用指针
	change(book1)
	fmt.Printf("%v\n", book1) // {1 《java并发编程的艺术》}
	changeBook(&book1)
	fmt.Printf("%v\n", book1) // {333 《java并发编程的艺术》}
}

func changeBook(b *Book) {
	b.no = 333
}

func change(book Book) {
	book.no = 666
}
