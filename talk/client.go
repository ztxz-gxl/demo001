package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Failed to dial net , err ===== ", err)
	}
	defer conn.Close()
	fmt.Println("请输入你要发送的内容！！")

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read string, err ===", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出。。。。")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Failed to write conn, err ====", err)
		}
	}
}
