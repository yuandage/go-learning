package main

import "fmt"
type Student struct {
	id   int
	name string
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student{1, "mike"}
	//类型查询,类型断言
	for index, data := range i {
		//第一个返回的是值,第二个返回判断结果
		if value, ok := data.(int); ok == true {
			fmt.Printf("i[%d]类型为int,值为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("i[%d]类型为string,值为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("i[%d]类型为Student,值为id=%d,name=%s\n", index, value.id, value.name)
		}
	}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d]类型为int,值为%d\n", index, value)
		case string:
			fmt.Printf("i[%d]类型为string,值为%s\n", index, value)
		case Student:
			fmt.Printf("i[%d]类型为Student,值为id=%d,name=%s\n", index, value.id, value.name)
		}
	}
}
