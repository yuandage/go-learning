package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	//连接服务器
	fmt.Println("输入聊天服务器IP和端口(默认使用 127.0.0.1:8000 回车确认)")
	var addr string
	fmt.Scanln(&addr)
	if len(addr)==0 {
		addr = "127.0.0.1:8000"
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("net.Dial err=", err)
		time.Sleep(2*time.Second)
		return
	}
	disconnectFlag := make(chan bool)
	fmt.Println("Connect successful!")
	fmt.Println("Enter -h to see help")

	defer conn.Close()

	//给服务器发送数据
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("err=", err)
				return
			}
			conn.Write(str[:n])
		}
	}()

	//接受服务器回复的数据
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					disconnectFlag <- true
					return
				}
				disconnectFlag <- true
				fmt.Println("err=", err)
				return
			}
			fmt.Println(string(buf[:n-1]))
		}
	}()

	for {
		select {
		case <-disconnectFlag:
			fmt.Println("Disconnect!")
			return
		}
	}
}
