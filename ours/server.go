package main

import (
	"fmt"
	"net"
	"strings"
)

var ConnMap = make(map[string]net.Conn)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("server start error")
	}

	defer listen.Close()
	fmt.Println("server is waiting ....")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn fail ...")
		}
		fmt.Println(conn.RemoteAddr(), "connect succeed")

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		data := make([]byte, 255)
		msgRead, err := conn.Read(data)
		if msgRead == 0 || err != nil {
			continue
		}

		//解析协议
		msgStr := strings.Split(string(data[0:msgRead]), "|")

		switch msgStr[0] {
		case "nick":
			fmt.Println(conn.RemoteAddr(), "-->", msgStr[1])
			for k, v := range ConnMap {
				if k != msgStr[1] {
					v.Write([]byte("[" + msgStr[1] + "]: join..."))
				}
			}
			ConnMap[msgStr[1]] = conn
		case "say":
			for k, v := range ConnMap {
				if k != msgStr[1] {
					fmt.Println("Send "+msgStr[2]+" to ", k)
					v.Write([]byte("[" + msgStr[1] + "]: " + msgStr[2]))
				}
			}
		case "quit":
			for k, v := range ConnMap {
				if k != msgStr[1] {
					v.Write([]byte("[" + msgStr[1] + "]: quit"))
				}
			}
			delete(ConnMap, msgStr[1])
		}

	}

}
