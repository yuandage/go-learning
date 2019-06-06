package main

import (
	"fmt"
	"time"
)

func main() {
	//延时功能,有三种方法
	//第1种
	//创建定时器 2s后向timer通道写数据
	timer := time.NewTimer(2 * time.Second) //时间到后,只会响应一次
	//timer.Stop()	//定时器停止
	//timer.Reset(4*time.Second)	//定时器重置
	fmt.Println("当前时间:", time.Now())

	//2s后
	t := <-timer.C
	fmt.Println("当前时间:", t)
	fmt.Println("2s时间到")

	//第2种
	<-time.After(2 * time.Second) //定时2s,阻塞2s,2s后向channel写数据
	fmt.Println("2s时间到")

	//第3种
	time.Sleep(time.Second * 2)
	fmt.Println("2s时间到")

	//时间到后,循环响应多次
	time1:=time.NewTicker(1*time.Second)
	i:=0
	for  {
		<-time1.C
		i++
		fmt.Println("i=",i)
		if i==5{
			time1.Stop()
			break
		}
	}
}
