package main

import "fmt"

func Test() int  {//无参有返回值
	return 666
}

func Test1() (result int)  {//无参有返回值 go推荐写法
	return 233
}

func Test3() (result int)  {//无参有返回值 go推荐写法
	return 110
}

func Test2() (a,b,c int)  {//无参有多个返回值
	a,b,c=233,666,111
	return
}

type TestType func() int //函数类型 可以实现多态,回调函数

func main() {
	var a int
	a=Test()
	fmt.Println(a)
	a=Test1()
	fmt.Println(a)

	b,c,d:=Test2()
	fmt.Println(b,c,d)

	var test TestType //多态
	test=Test1
	a=test()
	fmt.Println(a)

	test=Test3
	a=test()
	fmt.Println(a)
}