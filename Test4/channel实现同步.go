package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	go Person1()
	go Person2()
	for {
	}
}

func Person1() {
	Printer("hello")
	ch <- 666	//向通道写数据

}

func Person2() {
	<-ch	//从通道取数据,如果没有数据就会阻塞
	Printer("world")
}

func Printer(s string) {
	for _, data := range s {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}
