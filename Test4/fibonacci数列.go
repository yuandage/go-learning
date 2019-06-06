package main

import (
	"fmt"
)

func main() {
	ch:=make(chan int)	//数字通信
	quit:=make(chan bool)

	//消费者,读数列
	go func() {
		for i:=0;i<10 ;i++  {
			num:=<-ch
			fmt.Println("num=",num)
		}
		//读完后,可以停止
		quit<-true
	}()
	//生产者,产生数列
	fibonacci(ch,quit)
}

func fibonacci(ch chan<- int, quit <-chan bool) {
	x,y:=1,1
	for  {
		select {
		case ch<-x:
			x,y=y,x+y
		case flag:=<-quit:
			fmt.Println("flag=",flag)
			return
		}
	}
}
