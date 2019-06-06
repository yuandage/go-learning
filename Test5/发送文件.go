package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err=", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕!")
				return
			}
			fmt.Println("f.Read err=", err)
			return
		}
		conn.Write(buf[:n]) //发送文件内容
	}
}

func main() {
	fmt.Println("请输入需要传输的文件:")
	var path string
	fmt.Scan(&path)
	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err=", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err1=", err1)
		return
	}

	defer conn.Close()

	//先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//接收回复,ok表示可以可以发送文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	if string(buf[:n]) == "ok" {
		SendFile(path, conn) //发送文件内容
	}
}

