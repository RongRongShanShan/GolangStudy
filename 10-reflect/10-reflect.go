package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/**
reflect反射，每个变量都可以看做为pair<type,value>,能够通过变量反射找到type和value
	interface及其pair的存在，是Golang中实现反射的前提
结构体Tag标签及其在Json中的使用
*/

type User struct {
	Id   int    `json:"id"` // ``中的内容就是字段的标签，这里3个字段都设置了json标签
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) Call() {
	fmt.Printf("user is called ...  %v\n", u)
}

func main() {
	//知识点1，pair<type, value>
	// pair<statictype:string, value:"abc">
	var a string = "abc"
	// pair<type:string, value:"abc">
	var allType interface{} = a
	//通过类型断言取出值
	value, _ := allType.(string)
	fmt.Println(value)

	//知识点二，reflect包(反射包)获取变量的值和类型 (TypeOf和ValueOf)
	fmt.Println("value =", reflect.ValueOf(a), "type =", reflect.TypeOf(a))

	//知识点三，利用反射获取对象的字段和方法
	user := User{Id: 1, Name: "AAA", Age: 18}
	doFiledAndMethod(user)

	//知识点四，编码，结构体转json
	u := User{Id: 2, Name: "a", Age: 15}
	marshal, err := json.Marshal(u)
	if err != nil {
		println("json解析出错")
	} else {
		// 这里如果使用%v打印的结果是一个数组，以为json解析出的marshal的类型是byte[]
		fmt.Printf("jsonStr = %s\n", marshal)
	}

	//知识点五，解码，json转结构体
	jsonUser := User{}
	err = json.Unmarshal(marshal, &jsonUser)
	if err != nil {
		println("json解析出错")
	} else {
		fmt.Printf("jsonUser = %+v\n", jsonUser)
	}
}

// 利用反射通过type获取字段和方法
func doFiledAndMethod(arg interface{}) {
	//获取到type
	argType := reflect.TypeOf(arg)
	//获取到value
	argValue := reflect.ValueOf(arg)

	//通过type和value获取arg包含的字段类型和具体的值
	// 1. reflect.TypeOf(i any)得到reflect.Type; reflect.ValueOf(i arg)得到reflect.Value
	// 2. 通过Type的NumField获取字段数量，遍历
	// 3. 通过Type的Field(i int)得到StructField类型的第i+1个字段类型信息field
	// 4. 通过Value的Filed(i int)得到Value类型的对象，再通过Interface()得到第i+1个字段的值value
	fmt.Printf("输入对象包含以下对象\n")
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argValue.Field(i).Interface()
		fmt.Printf("  %s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历 (调用类型为*User时不会显示)
	for i := 0; i < argType.NumMethod(); i++ {
		m := argType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
