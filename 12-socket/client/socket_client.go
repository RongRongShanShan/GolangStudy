package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 使用tcp连接本地8080端口的客户端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	var message string
	for {
		// 阻塞写入信息
		_, err := fmt.Scan(&message)
		if err != nil {
			fmt.Printf("%+v", err)
			return
		}
		// 向服务端发送写入的信息
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}
		// 阻塞获取回复
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error receiving message:", err.Error())
			return
		}
		fmt.Println("Received message:", string(buf[:n]))
	}
}
