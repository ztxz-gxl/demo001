package main

import (
	"fmt"
	"io"
	"net" //socket 编程必备
)

func main() {
	fmt.Println("-------------服务器开始监听-------------------")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Failed to net.Listen,  err ===== ", err)
	}
	defer listen.Close()

	for {
		//		fmt.Println("等待客户端来链接。。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Fail to Accept conn,  err === ", err)
		}
		fmt.Println("客户端链接成功", conn.RemoteAddr().String())
		go Process(conn)
	}
}

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("远程客户端%v退出\n", conn.RemoteAddr())
				conn.Close()
				return
			}
			fmt.Println("Failed to read conn,  err == ", err)
			conn.Close()
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
