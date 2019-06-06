package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	disconnectFlag := false
	addr := conn.RemoteAddr().String()
	fmt.Printf("[%s]:Connect successful!\n", addr)

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
			fmt.Printf("[%s]:", addr)
			n, err := conn.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					disconnectFlag = true
					fmt.Printf("[%s]:Disconnect!\n", addr)
					return
				}
				fmt.Println("err=", err)
				return
			}
			fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		}
	}()

	for {
		if disconnectFlag == true {
			return
		}
	}
}
