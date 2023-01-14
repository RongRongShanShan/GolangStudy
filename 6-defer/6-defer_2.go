package main

/**
defer与return的先后
	defer在函数结束后才会被调用，所以在return之后
*/

func deferFunc() int {
	println("deferFunc 被调用")
	return 0
}

func returnFunc() int {
	println("returnFunc 被调用")
	return 1
}

func test() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	i := test()
	println(i)
}
