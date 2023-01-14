package main

import "fmt"

/**
slice，切片，动态数组，是对数组的抽象
*/

// 1. 先介绍数组 和 动态数组(slice)
func studyArray() {
	println("\n=== 数组和动态数组(slice)的介绍 ===\n")
	// 创建固定长度数组,未初始化的值都用默认值
	var a [5]int
	b := [4]int{1, 2, 3}
	// 对数组进行遍历 - 方式1
	// 	注意go的for没有括号，并且go只有后置++
	for i := 0; i < len(a); i++ {
		println(a[i])
	}
	// 对数组进行遍历 - 方式2
	// 	使用range关键字遍历索引和值 (只用一个参数接收时接收到第一个返回值(就是索引))
	for index, value := range b {
		println("index = ", index, "value = ", value)
	}
	// 查看数组类型
	fmt.Printf("type of a is %T\n", a) // type of a is [5]int
	fmt.Printf("type of b is %T\n", b) // type of b is [4]int

	// 使用这样固定长度的数组并不方便，因为函数参数传递默认是按值传递，如果要传递[5]int，函数形参也要为[5]int(进行值拷贝)
	// 	这样情况下[4]int的数组将无法作为该函数的参数，所以推荐使用动态数组[]int

	println("=== 动态数组 ===")
	// 动态数组，切片，slice (本质上是指向数组位置的指针)
	c := []int{1, 2, 3}
	// 动态数组类型
	fmt.Printf("type of c is %T\n", c) // type of c is []int
	// 通过函数遍历并修改动态数组，发现可以在函数中直接修改动态数组 (传递的是引用)
	// 动态数组在传递参数时就不需要考虑数组长度
	printArray(c)
	println(c[0]) //c[0]被修改成4
}

func printArray(arr []int) {
	// 下划线表示匿名对象
	for _, value := range arr {
		println("value = ", value)
	}
	// 修改动态数组中的元素，会直接修改原值
	arr[0] = 4
}

// 2. 切片使用方式
func useSlice() {
	println("\n=== 切片(slice/动态数组)的使用 ===\n")
	// 初始化slice 4种方式
	//slice1 := []int{1,2,3}  //声明切片并且初始化，值为[1 2 3]，长度为3
	var slice1 []int //声明切片，但没有分配空间  (注意 可以对空切片直接进行追加)
	//var slice1 []int = make([]int, 3) //声明切片，并通过make分配3个空间，初始值为元素类型的默认值(0)
	//slice1 := make([]int, 3) //make分配空间并通过:=推导出类型为切片

	// 打印切片数据和类型
	fmt.Printf("slice1 = %v, type of it is %T\n", slice1, slice1) // slice1 = [0 0 0], type of it is []int

	// 判断slice是否没有任何元素(没开辟空间)
	// nil表示零值/空值
	if slice1 == nil {
		println("slice1 是空切片")
	} else {
		println("slice1 分配了空间")
	}

	// 切片容量的追加  (重点)
	nums := make([]int, 3, 4)                                                  // 声明切片的长度为3，但容量为4 (容量是实际开辟空间的大小；如果没有声明cap，cap=len)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums, len(nums), cap(nums)) // slice = [0 0 0], len = 3, cap = 4
	// 向nums切片中追加一个元素1， len+1，cap不变，元素追加到末尾
	nums = append(nums, 1)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums, len(nums), cap(nums)) // slice = [0 0 0 1], len = 4, cap = 4
	// 向len==cap的切片继续追加元素， 开辟新空间以满足容量需求，cap翻倍，len+1元素追加到末尾
	nums = append(nums, 2)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums, len(nums), cap(nums)) // slice = [0 0 0 1 2], len = 5, cap = 8

	// 切片的截取  (重点)
	// go切片的截取与py的切片类似，形式是slice[start, end],截取的范围是[start,end)
	nums2 := nums[0:5]                                                            //从开始截取4个位置，[0 0 0 1]
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums2, len(nums2), cap(nums2)) // slice = [0 0 0 1], len = 4, cap = 8
	// 根据输出结果，可以看到cap依旧为8，这是因为nums2实际上使用的空间和num相同，只是len不一样，如果对nums2进行修改，nums相同位置也会被修改
	nums2[0] = 100
	println("nums[0] = ", nums[0], "; nums2[0] = ", nums2[0]) // nums[0] =  100 ; nums2[0] =  100
	// 如果截取是开始位置不为0，cap会相应减少
	nums3 := nums[1:6]
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums3, len(nums3), cap(nums3)) // slice = [0 0 1 2 0], len = 5, cap = 7
	// 如果要使截取的结果生成一个新数组，可以使用内置函数copy
	nums4 := make([]int, 5)
	copy(nums4, nums3)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", nums4, len(nums4), cap(nums4)) // slice = [0 0 1 2 0], len = 5, cap = 5
	// 对新数组的修改不会影响原数组
	nums4[0] = 200
	println("nums3[0] = ", nums3[0], "; nums4[0] = ", nums4[0]) // nums3[0] =  0 ; nums4[0] =  200
}

func main() {
	studyArray()
	useSlice()
}
