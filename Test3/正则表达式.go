package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac 43.14 567 agsdg 1.23 7. 8.9"
	//解析正则表达式,成功返回解释器
	reg := regexp.MustCompile(`a.c`) //`` 原生字符串
	reg = regexp.MustCompile(`a[0-9]c`)
	reg = regexp.MustCompile(`a\Dc`)     //\D非数字
	reg = regexp.MustCompile(`\d+\.\d+`) //+匹配前一个字符一次或多次

	//根据规则提取关键信息
	//result:=reg.FindAllString(buf,-1)	//不分组
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result)

	buf1 := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Document</title>
			</head>
			<body>
				<div>测试
				类比
				比喻
				</div>
				<div>666</div>
				<div>233</div>
				<div>你不要过来啊</div>
				
			</body>
			</html>`

	reg1 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)	//没有换行
	result1 := reg1.FindAllStringSubmatch(buf1, -1)
	fmt.Println("result=", result1)

	for _,text:=range result1{
		fmt.Println("text[0]=",text[0])	//带<></>
		fmt.Println("text[1]=",text[1])	//不带<></>
	}
}
