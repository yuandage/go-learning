package main

import "fmt"

//定义接口类型
type Humaner interface {	//子集
	SayHi()	//方法只有声明,没有实现
}

//接口继承
type Personer interface {	//超集
	Humaner
	Sing()
}

//实现接口
func (p Person) SayHi()  {
	fmt.Println("Person SayHi")
}

func (p Student1) SayHi()  {
	fmt.Println("Student1 SayHi")
}

func (p Student1) Sing()  {
	fmt.Println("Student1 Sing")
}

//定义一个普通函数,参数为接口,体现多态
func WhoSayHi(i Humaner)  {
	i.SayHi()
}

type Person struct {
	name string
	sex  byte
	age  int
}

type Student1 struct {
	Person //只有类型的匿名字段 继承Person的成员
	id   int
	addr string
	name string //就近原则,首先寻找本作用域,找不到就找继承的
}

type Student2 struct {
	*Person //只有类型的匿名字段 继承Person的成员
	id   int
	addr string
	name string //就近原则,首先寻找本作用域,找不到就找继承的
}

type add int //方法的使用
//tmp 叫接收者,就是传递的一个参数
func (tmp add) Add(other add) add {
	return tmp + other
}

//SetInfo方法,给成员赋值
func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

//GetInfo方法,获取成员值
func (p Person) GetInfo() {
	fmt.Println("Person=", p)
}

func (p *Person) GetInfo1() {
	fmt.Println("Person=", *p)
}

//GetInfo方法重写
func (p Student2) GetInfo() {
	fmt.Println("Student2=", p)
}

func main() {
	var s1 Student1 = Student1{Person{"mike", 'm', 18}, 1, "bj", "mike666"}
	s1.Person.name = "mike233"
	fmt.Printf("%+v\n", s1) //%+v 详细打印
	fmt.Println("name:", s1.name, "age:", s1.age)
	//初始化方式一
	var s2 Student2 = Student2{&Person{"mike", 'm', 18}, 1, "bj", "mike666"}
	fmt.Printf("%+v\n", s2) //%+v 详细打印
	fmt.Println(s2.Person, s2.id, s2.addr, s2.name)
	//初始化方式二
	var s3 Student2
	s3.Person = new(Person)
	s3.GetInfo()	//方法继承

	var result add = 2
	result = result.Add(3)
	fmt.Println(result)

	var p Person
	(&p).SetInfo("john", 'm', 18)
	p.GetInfo()

	f:=p.GetInfo	//方法值,保存了方法的入口地址,隐藏了接收者
	f()

	f1:= (*Person).GetInfo1 //方法表达式
	f1(&p)

	var i Humaner	//定义接口类型变量
	s4:=Student1{Person{"mike",'m',18},1,"bj","666"}
	p1:=Person{"tony",'m',18}
	//调用接口
	i=s4
	i.SayHi()
	WhoSayHi(i)

	i=p1
	i.SayHi()
	WhoSayHi(i)

	var i1 Personer
	i1=s4
	i1.SayHi()	//调用继承的接口
	i1.Sing()

	//超集可以转换为子集,子集不能转换为超集
	i=i1
	i.SayHi()

	var i2 interface{}="666"	//空接口类型,保存任意类型,相当于Object类型
	fmt.Println(i2)	//func Println(args ...interface{}),函数可以接收任意类型

}
