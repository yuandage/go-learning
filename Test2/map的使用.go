package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println("m1=",m1)
	fmt.Println("len(m1)=",len(m1))	//map没有cap()

	m2:=make(map[int]string,10)
	fmt.Println("m2=",m2)
	fmt.Println("len(m2)=",len(m2))

	m3:=map[int]string{1:"go",2:"c++",3:"java"}	//键值唯一
	fmt.Println("m3=",m3)
	fmt.Println("len(m3)=",len(m3))
	m3[4]="python"
	fmt.Println("m3=",m3)

	for key,value:=range m3 {
		fmt.Println(key,value)
	}
	value,ok:=m3[0]
	if ok==true{
		fmt.Println("m3[1]=",value)
	}else{
		fmt.Println("key不存在")
	}
	delete(m3,3)
	fmt.Println("m3=",m3)
}