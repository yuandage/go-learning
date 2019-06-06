package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	listenr, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer listenr.Close()
	//阻塞等待客户端连接
	for {
		conn, err := listenr.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//接收客户端请求,新建一个协程
		go HandleConn(conn)

	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close() //关闭客户端连接
	addr := conn.RemoteAddr().String()
	fmt.Printf("[%s]:Connect successful!\n", addr)

	buf := make([]byte, 2048)
	//读取客户端数据
	for {

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		if string(buf[:n-1]) == "quit" || string(buf[:n-2]) == "quit" {
			fmt.Printf("[%s]:Disconnect!\n", addr)
			return
		}

		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		//将读到的数据转换为大写,再发送给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
