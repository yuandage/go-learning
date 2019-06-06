package main

import (
	"fmt"
	"time"
)

func main() {
	//channel 默认双向
	//ch:=make(chan string)	//无缓冲channel
	//通道写完,没有读会阻塞,通道没有数据,读数据会阻塞

	//有缓冲channel
	ch := make(chan int, 3)
	fmt.Println("len(ch)=", len(ch), "cap(ch)", cap(ch))
	defer fmt.Println("主协程结束")

	go func() {
		defer fmt.Println("子协程结束")

		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("i=", i)
			time.Sleep(time.Second)
		}
		close(ch) //关闭通道后,不能再向channel写数据
	}()

	//延时
	time.Sleep(2 * time.Second)
	//for {
	//	if num, ok := <-ch; ok == true { //读通道内容,没有内容,阻塞
	//		fmt.Println("num=", num)
	//	} else {
	//		break
	//	}
	//}
	//简便写法
	for num := range ch {
		fmt.Println("num=", num)
	}

	//ch1 := make(chan int)
	////双向channel能隐式转换为单向channel
	//var writeCh chan<- int = ch1 //只能写,不能读
	//var readCh <-chan int = ch1  //只能读,不能写
	//
	//writeCh <- 666
	//<-readCh
	//单向channel不能转换为双向

	//生产者,消费者
	ch2:=make(chan int)
	go producer(ch2)
	consumer(ch2)
}

func producer(out chan<- int) {
	for i:=0;i<10 ;i++  {
		out<-i*i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num:=range in{
		fmt.Println("num=",num)
	}
}
