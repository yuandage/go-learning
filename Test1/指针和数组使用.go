package main

import "fmt"

func main() {
	a := 10
	var p *int
	p = &a
	fmt.Println(*p)

	p1 := new(int)
	*p1 = 666
	fmt.Println(*p1)

	var id [10]int //定义数组
	fmt.Println(id)

	for i := 0; i < len(id); i++ {
		id[i] = i + 1
	}

	for i, data := range id {
		fmt.Printf("id[%d]=%d\n", i, data)
	}

	//var b1 [5]int  =[5]int{1,1,1,1,1}
	var b1 = [5]int{1, 1, 1, 1, 1}
	fmt.Println(b1)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c:=[5]int{1,2,3}
	d:=[5]int{2:10,4:20}
	fmt.Println(c,d)

	e:=[3][3]int{1:{1,2,3}}
	fmt.Println(e)
}
