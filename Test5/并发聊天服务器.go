package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送的管道
	Name string      //用户名
	Addr string      //ip地址
}

var onlineMap map[string]Client
var message = make(chan string)

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息时,会阻塞
		//给每个用户发送消息
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	//获取用户ip地址
	clientAddr := conn.RemoteAddr().String()
	client := Client{make(chan string), clientAddr, clientAddr}
	onlineMap[clientAddr] = client

	//新协程,给当前用户发送消息
	go WriteToClient(client, conn)

	//广播某个用户在线
	message <- MakeMsg(client, "login")
	//提示用户地址
	client.C <- MakeMsg(client, "I am here !")
	//新协程,接收用户的消息
	isQuit := make(chan bool)  //用户主动退出
	hasData := make(chan bool) //超时退出

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //连接断开,或出错
				isQuit <- true
				fmt.Println("conn.Read err", err)
				return
			}
			msg := string(buf[:n-1]) //"\r\n"
			if string(buf[:n-1])=="-h" || string(buf[:n-2]) == "-h" {
				msg = "1. -h :see help\n" +
					"2. -user :see user information\n" +
					"3. rename [username] :change username\n" +
					"4. -quit :quit client\n"
				conn.Write([]byte(msg))
			} else if string(buf[:n-1])=="-user" || string(buf[:n-2]) == "-user" {
				//遍历map,发送所有用户信息
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, " ")[1]
				client.Name = name
				onlineMap[clientAddr] = client
				conn.Write([]byte("rename ok !\n"))
			} else if string(buf[:n-1])=="-quit" || string(buf[:n-2]) == "-quit" {
				isQuit <- true
			} else {
				//转发内容
				message <- MakeMsg(client, msg)
			}
			hasData <- true //有数据
		}

	}()

	for {
		select {
		case <-isQuit:
			message <- MakeMsg(client, "logout") //广播下线
			delete(onlineMap, clientAddr)
			return
		case <-hasData:
		case <-time.After(60 * time.Second): //60s 超时下线
			message <- MakeMsg(client, "time out leave")
			delete(onlineMap, clientAddr)
			return
		}
	}
}

func MakeMsg(client Client, msg string) string {
	return "[" + client.Addr + "]" + client.Name + ":" + msg
}

func WriteToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	fmt.Println("TCP并发聊天服务器运行中...\nIP:127.0.0.1:8000")
	defer listener.Close()

	//新协程,转发消息
	go Manager()

	//主协程,循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err=", err)
			continue
		}

		go HandleConnection(conn) //处理用户连接
	}

}
