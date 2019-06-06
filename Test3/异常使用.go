package main

import (
	"fmt"
)
import "errors"

func Div(a,b int) (result int,err error) {
	err=nil
	if b==0 {
		err=errors.New("被除数不能为0")
	}else{
		result=a/b
	}
	return
}

func test() {
	panic("this is panic test")	//显示调用panic函数,使程序中断
}

func test1(x int)  {
	//recover()	可以打印panic的错误信息
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()

	var a [10]int
	a[x]=111
}

func main() {
	result,err:=Div(10,0)
	if err!=nil{
		fmt.Println("err=",err)
	}else{
		fmt.Println("result=",result)
	}

	//test()
	test1(20)

}