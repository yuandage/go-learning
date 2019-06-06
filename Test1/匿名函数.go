package main

import "fmt"

func testbibao() func() int {
	var x int

	return func() int { //形成闭包 只要有闭包在使用变量,变量就会一直存在
		x++
		return x*x
	}
}

func main() {
	a := 10
	str := "mike"

	//闭包以引用的方式捕获变量
	f1 := func() { //匿名函数 函数定义
		fmt.Println("a=", a)	//形成闭包 可以捕获外面的变量
		fmt.Println("str=", str)	//可以修改外面变量的值
	}

	f1() //函数调用

	type FuncType = func() //给匿名函数起别名 很少使用
	var f2 FuncType
	f2 = f1
	f2()

	func() {//定义匿名函数,同时调用
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}() //()调用匿名函数

	f3:=func(i,j int) {//带参数的匿名函数
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}
	f3(1,2)

	func(i,j int) {//带参数的匿名函数
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}(1,2)	//定义并调用匿名函数

	//匿名函数有参有返回值
	x,y:= func(i,j int) (max,min int) {
		if i>j{
			max=i
			min=j
		}else{
			max=j
			min=i
		}
		return
	}(10,20)
	fmt.Println(x,y)

	f4:=testbibao()
	fmt.Println(f4())
	fmt.Println(f4())
	fmt.Println(f4())
	fmt.Println(f4())

	//defer 延迟调用,mian函数结束前调用
	defer fmt.Println("111")	//当程序出错,defer后的语句都会调用
	defer fmt.Println("222")
	defer fmt.Println("333")
	fmt.Println("bbb")


}
