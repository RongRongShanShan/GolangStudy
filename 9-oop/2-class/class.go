package main

import "fmt"

/**
go的类方法还是在结构体的基础上进行的
	这里演示包括go的继承和多态，严格来说go没有继承和多态; 所谓的继承其实是通过匿名组合实现的伪继承，多态则是结构体隐式地实现接口
	还包括空接口interface{}作万能数据类型及其类型断言机制
*/

// Book 定义结构体
//
//	注意！ 类名和其方法首字母大写意味着其他包也可以访问到，而首字母小写(比如Book中定义的两个对象)则视为私有
//		跨包只能调用其他包中对外开放的方法，(类名、属性名。方法名)首字母大写代表对其他包开放，go语言的封装是针对包来说的，在同一包内大小写没有区别
type Book struct {
	no   int
	name string
}

// Show 声明Book的方法，func后面的()中指明了调用的类型为Book，使用*Book是为了通过引用操作原数据而不是生成的副本(struct不使用引用时会操作生成的副本)
func (b *Book) Show() {
	fmt.Printf("%v\n", *b)
}

func (b *Book) SetNo(no int) {
	b.no = no
}

func (b *Book) SetName(name string) {
	b.name = name
}

func main() {
	//类方法
	book := Book{no: 1, name: "golang入门"}
	book.Show() // {1 golang入门}
	book.SetName("golang面向对象")
	book.Show() // {1 golang面向对象}

	fmt.Println("=== 继承 ===")
	//继承(伪继承)
	people1 := People{name: "张三"}
	people1.show()
	superman := SuperMan{
		People: People{"超人"},
		no:     1,
	}
	superman.show()

	fmt.Println("=== 多态 ===")
	//多态 隐式实现接口中声明的方法的类可以用接口接收和使用
	var a1 animal = &cat{name: "1"}
	var a2 animal = &dog{name: "1"}
	a1.say()
	a1.sleep()
	a2.say()
	a2.sleep()

	fmt.Println("=== 万能数据类型 ===")
	//任何类型(包括int等)都可以看作隐式实现了空接口interface{},因为它没有声明任何方法
	//	我们可以以interface{}作为形参的类型接收任何类型的参数
	showAnything(1)                                         // 变量内容：1 , 类型：int  ...
	showAnything(3.4)                                       // 变量内容：3.4 , 类型：float64  ...
	showAnything(map[string]string{"a": "aaa", "b": "bbb"}) // 变量内容：map[a:aaa b:bbb] , 类型：map[string]string  ...
	showAnything(superman)                                  // 变量内容：{People:{name:超人} no:1} , 类型：main.SuperMan  ...
	showAnything("a")
}

//下面演示继承(严格来说go没有继承，是通过”匿名组合”来实现继承的效果)

// People 父类
type People struct {
	name string
}

// 父类方法，展示内容
func (p *People) show() {
	fmt.Printf("%+v\n", *p)
}

type SuperMan struct {
	//继承
	People // 相当于写了个 people People，是一个匿名对象，组合的思想

	no int
}

// 重写父类方法
func (m *SuperMan) show() {
	fmt.Printf("%+v\n", *m)
}

//下面演示多态

// 定义接口声明方法
type animal interface {
	say()
	sleep()
}

// cat 实现接口是隐式实现的，需要实现接口声明的所有方法
type cat struct {
	name string
}

// 直接实现接口声明的所有方法就可以使用接口接收该类对象，实现多态
func (c *cat) say() {
	fmt.Printf("猫猫叫\n")
}
func (c *cat) sleep() {
	fmt.Printf("猫猫睡觉\n")
}

// dog 在定义一个实现接口的类
type dog struct {
	name string
}

func (d *dog) say() {
	fmt.Printf("狗狗叫\n")
}
func (d *dog) sleep() {
	fmt.Printf("狗狗睡觉\n")
}

//下面展示空接口interface{}作万能数据类型

func showAnything(arg interface{}) {
	// interface{} 提供有 类型断言机制； 如下判断是否为string类型
	value, ok := arg.(string)
	if !ok {
		// 根据多态会展示作为参数的对象原本的数据
		fmt.Printf("内容：%+v , 类型：%T\n", arg, arg)
	} else {
		fmt.Printf("%s是string类型", value)
	}
}
