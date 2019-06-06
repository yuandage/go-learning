package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func test1(args ...int){//不定参数,放在参数列表最后
	fmt.Println("len(args)=",len(args))
	for i, data := range args {	//返回下标和元素值
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}

func test2(args ...int){//不定参数,放在参数列表最后
	fmt.Println("len(args)=",len(args))
	for i, data := range args {	//返回下标和元素值
		fmt.Printf("args[%d]=%d\n", i, data)
	}
	test1(args[:2] ...)//传入形参args[0]-args[1]
	test1(args[2:] ...)//传入形参args[2]-args[n]
}

func main() {
	var a string
	a = "变量"
	var b int = 10;
	fmt.Println("hello go" + a)
	fmt.Println(b)
	fmt.Printf("b的类型 %T\n", b)

	i, j := 10, 20

	i, j = j, i
	fmt.Printf("i=%d,j=%d\n", i, j)

	var c, d, e int
	c, d, e = test()
	fmt.Printf("c %d d %d e %d\n", c, d, e)

	_, d, _ = test()
	fmt.Printf("d %d \n", d) //匿名变量和和函数返回值配合使用
	const k = 10
	var (
		m int
		l float64
	)
	var (
		m1 = 10
		l1 = 3.14
	)
	fmt.Println("自动推导类型", m1, l1)
	m = 10
	l = 3.14
	fmt.Println(m, l)

	const (
		a1 int     = 10
		a2 float64 = 3.14
	)
	const (
		a3 = 10
		a4 = 3.14
	)
	fmt.Println("自动推导类型", a3, a4)

	str := "mike"
	fmt.Println(str, len(str))
	fmt.Printf("%c\n", str[0])

	//fmt.Println("输入字符串:")
	//fmt.Scan(&str)
	//fmt.Println("str:",str)

	if a := 10; a == 10 {
		fmt.Println("a=10")
	}

	switch a := 3; a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3, 4, 5:
		fmt.Println(3, 4, 5)
	default:
		fmt.Println("其他")
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)
	//range 迭代
	str1 := "abc"
	for i, data := range str1 {	//返回下标和元素值
		fmt.Printf("str1[%d]=%c\n", i, data)
	}

	test1(1,2,3)

}
