package main

import "fmt"

func Add(a,b int) int {
	return a+b
}

func Sub(a,b int) int {
	return a-b
}

func Mul(a,b int) int {
	return a*b
}

func Div(a,b int) int {
	return a/b
}

type fTest func(a,b int) int //函数类型 可以实现多态,回调函数

func Calc(a,b int, option fTest )(result int){
	result=option(a,b)
	return
}

func main() {
	a:=Calc(5,5,Add)
	fmt.Println(a)

	a=Calc(5,5,Sub)
	fmt.Println(a)

	a=Calc(5,5,Mul)
	fmt.Println(a)

	a=Calc(5,5,Div)
	fmt.Println(a)
}