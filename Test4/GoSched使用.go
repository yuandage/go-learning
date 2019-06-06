package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i:=0;i<5;i++ {
			fmt.Println("go i=",i)
			if i==4{
				runtime.Goexit()	//终止所在协程
			}
		}
	}()

	for i:=0;i<2 ;i++  {
		runtime.Gosched()	//让出时间片,先让其他协程执行,再回来执行此协程
		fmt.Println("hello i=",i)
	}
}
