package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf) //接收文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕!")
				return
			}
			fmt.Println("conn.Read err=", err)
			return
		}
		f.Write(buf[:n]) //写入文件内容
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listener.Close()

	//等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1=", err1)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf) //接收文件名
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	fileName := string(buf[:n])
	//回复ok
	conn.Write([]byte("ok"))
	//接收文件内容
	RecvFile(fileName, conn)
}


