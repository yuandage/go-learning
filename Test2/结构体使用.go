package main

import "fmt"

type Student struct {	//结构体
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	s1.addr = "bj123"
	fmt.Println(s1)

	s2 := Student{name: "tony", addr: "cd"}
	fmt.Println(s2)

	s3 := &Student{name: "tony", addr: "cd"}
	fmt.Println(*s3, s3)

	s4 := new(Student) //通过new初始化一个结构体
	s4.name = "john"
	fmt.Println(s4)

}
