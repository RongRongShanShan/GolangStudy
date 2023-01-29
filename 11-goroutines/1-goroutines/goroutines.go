package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
协程goroutine，相当于轻量级的线程
	Sync并发包
*/

// HelloGoRoutine 通过协程快速打印
func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		//go关键字:开启协程
		go func(j int) {
			println("hello goroutine : " + fmt.Sprint(j))
		}(i) // 最后的括号表示函数自调用，括号中的参数是形参
	}
	//避免主线程先结束
	time.Sleep(time.Second)
}

// ManyGoWait 通过计数器WaitGroup对HelloGoRoutine()方法进行优化
func ManyGoWait() {
	var wg sync.WaitGroup
	//为计数器设置初始值
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			//每个任务完成时将计数器-1
			defer wg.Done()
			println("hello goroutine : " + fmt.Sprint(j))
		}(i)
	}
	//并发任务完成前等待(计数器不为0时阻塞)
	wg.Wait()
}

func main() {
	//HelloGoRoutine()

	//一个演示
	go func() {
		defer fmt.Println("A defer")
		func() {
			defer fmt.Println("B defer")
			//这里退出了线程
			runtime.Goexit()
			fmt.Println("B")
		}() //调用定义的匿名函数
		fmt.Println("A")
	}() // 结果只有两个defer被打印
	//主线程睡眠1s保证协程执行
	time.Sleep(1 * time.Second)
}
