package main

import (
	"fmt"
	"time"
)

func main() { //主协程退出后,子协程也会退出

	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程i=", i)
			time.Sleep(time.Second)
		}
	}()

	go newTask() //新建一个协程

	i:=0
	for {
		i++
		fmt.Println("main i=", i)
		time.Sleep(time.Second)
		if i==2{
			break	//主协程退出后,子协程也会退出
		}
	}

}

func newTask() {
	for {
		fmt.Println("this is a newTask routine")
		time.Sleep(time.Second)
	}
}
