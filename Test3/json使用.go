package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"-"` //二次编码,此字段不会输出到屏幕
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"` //转换为字符串
	Price    float64  `json:",string"`
}

type JSON struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	//buf,err:=json.Marshal(s)	//编码,根据结构体生成json文本
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))

	fmt.Println("------")

	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	buf, err = json.MarshalIndent(m, "", "	") //格式化编码
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))

	jsonBuf := `{
				"company":"itcast",
				 "subjects": [
				  "Go",
				  "C++",
				  "Python",
				  "Test"
				 ],
				 "isok": true,
				 "price": 666.666
				}`
	var tmp JSON
	//json解析到结构体
	err=json.Unmarshal([]byte(jsonBuf),&tmp)	//第二个参数为地址传递
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n",tmp)

	m1 := make(map[string]interface{}, 4)
	//json解析到map
	err=json.Unmarshal([]byte(jsonBuf),&m1)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n",m1)
	var str string
	//通过类型断言,来取出map里的值
	for key,value:=range m1 {
		switch data:=value.(type) {
		case string:
			str=data
			fmt.Printf("m[%s]=%s\n",key,str)
		case []interface{}:
			fmt.Printf("m[%s]=%v\n",key,data)
		}
	}
}
