package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//fmt.Println(rand.Int())
	//fmt.Println(rand.Intn(100)) //100以内的随机数

	var a [10]int
	n := len(a)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)

	}
	fmt.Println(a)

	fmt.Println("冒泡排序:")

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Println(a)

	b:= []int{1,2,3,0,0}	//切片---slice
	s:=b[0:3:5]	//[low:high:max]---切片
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s))
	fmt.Println("cap(s)=",cap(s))

	var c []int
	fmt.Println("len(c)=", len(c),"cap(c)=", cap(c))
	c=append(c, 11)	//给切片追加一个元素
	fmt.Println("append: len(c)=", len(c),"cap(c)=", cap(c))

	s1:=make([]int,5,10)	//切片创建 make(切片类型,长度,容量)
	fmt.Println(s1, len(s1), cap(s1))
	s2:=make([]int,5)
	fmt.Println(s2,len(s2), cap(s2))

	s1=s1[:]	//等价于s1[0,len(s1):len(s1)]




}
