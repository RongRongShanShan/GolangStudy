package main

import (
	"fmt"
	"time"
)

/**
通道channel，可以被认为是Goroutines通信的管道。类似于管道中的水从一端到另一端的流动，数据可以从一端发送到另一端，通过通道接收。
	go中 “不要通过共享内存来通信，而应该通过通信来共享内存”
*/

// CalSquare 协程通信(通过通道channel) : A子协程发送0-9数字，B子协程计算数字的平方，主协程输出 (生产消费模型)
func CalSquare() {
	//定义channel make(chan Type, caption)
	src := make(chan int)     //无缓冲的channel
	dest := make(chan int, 3) //带缓冲的channel
	//A协程
	go func() {
		//延迟src的关闭
		defer close(src)
		//0-9加入src
		for i := 0; i < 10; i++ {
			//将i放入src
			src <- i
		}
	}()
	//B协程
	go func() {
		defer close(dest)
		//获取src中的值，平方后加入dest
		for i := range src {
			dest <- i * i
			// 看一下有缓冲的channel在运行过程中的情况
			fmt.Println("dest len = ", len(dest), "cap = ", cap(dest))
		}
	}()

	//打印结果
	/*for i := 0; i < 10; i++ {
		num := <-dest
		fmt.Println(num)
	}*/
	for i := range dest {
		println(i)
	}
}

func selectAndChannel() {
	c := make(chan int)
	quit := make(chan int)
	defer close(c)
	defer close(quit)

	go func() {
		for i := 0; i < 6; i++ {
			//c可读时读出打印
			fmt.Println(<-c)
		}
		//对c写入5次后向quit写入，quit变得可读，退出
		quit <- 0
	}()

	//下面是斐波那契的逻辑
	x, y := 1, 1
	for {
		select {
		case c <- x: //如果c可写
			x, y = y, x+y
		case <-quit: //如果quit可读
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	/*CalSquare()
	time.Sleep(time.Second)*/
	selectAndChannel()
	time.Sleep(time.Second)
}
