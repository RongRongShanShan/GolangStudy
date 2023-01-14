package main

import "fmt"

/**
map，<key,value>形式的键值对，和java中Map类似，初始化和slice一样要使用make
*/

func main() {
	println("\n=== map声明方式 ===\n")

	// 第一种声明  类型格式是 map[key]value
	var map1 map[string]int
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	map1 = make(map[string]int, 10)
	// 添加和修改map元素直接 map[key] = value
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Println(map1) // map[one:1 three:3 two:2]  (注意 key没按插入顺序排序，依据哈希排列)

	// 第二种声明  变量名 := make(map[key_type]value_type, ...)
	map2 := make(map[int]string)
	map2[1] = "one"
	map2[2] = "two"
	map2[3] = "three"
	fmt.Println(map2) // map[1:one 2:two 3:three]

	// 第三种声明  声明时直接初始化  变量名 := map[key_type]value_type{ key1:value1,key2:value2... }
	map3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(map3) //map[one:php two:golang three:java]

	println("\n=== map基本使用 ===\n")

	// 添加
	map1["four"] = 4
	// 遍历
	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}
	// 删除 (按key删)
	delete(map1, "one")
	// 修改 同添加
	map1["four"] = 5
	// map和slice一样可以在其他函数中修改，因为本质是引用

	// 嵌套 (和其他语言差不多)
	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"
	fmt.Println(language) //map[php:map[id:1 desc:php是世界上最美的语言] golang:map[id:2 desc:golang抗并发非常good]]
}
