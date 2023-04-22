package main

import (
	"fmt"
	"net"
)

func main() {
	// 使用tcp协议监听本地的8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	// defer保证最后关闭监听
	defer func(listener net.Listener) {

		if err := listener.Close(); err != nil {
			fmt.Printf("%+v", err)
		}
	}(listener)

	// 循环监听是否有客户端的连接请求
	for {
		// 接受客户端的连接请求 (会阻塞)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("%+v", err.Error())
			continue
		}

		// 对每一个连接请求单独开一个协程处理
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	// 关闭资源
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			fmt.Printf("%+v", err.Error())
		}
	}(conn)

	// 不断阻塞监听客户端发送的消息
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Received message from %s : %s\n", conn.RemoteAddr().String(), string(buf))

		if _, err = conn.Write([]byte("ok")); err != nil {
			return
		}
	}
}
