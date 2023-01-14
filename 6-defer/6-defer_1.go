package main

/**
有关defer语句的调用顺序
	defer用于预定对一个函数的调用,在当前函数执行结束之后才调用
	如果一个函数中有多个defer语句，它们会以LIFO(后进先出)的顺序执行 (本质上就是栈)

defer的作用：
	● 释放占用的资源
	● 捕捉处理异常
	● 输出日志
*/

/*func main() {
	defer println("end-1")
	defer println("end-2")
	println("start-1")
	println("start-2")
}*/

/**
执行结果：
	start-1
	start-2
	end-2
	end-1
*/
