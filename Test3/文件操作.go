package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//标准设备文件(os.Stdout),默认打开,用户可以直接使用
	//os.Stdout.Close()	//关闭后,无法输出
	fmt.Println("are u ok ?")
	os.Stdout.WriteString("thank u very much\n")

	//os.Stdin.Close()	//关闭后,无法输入
	//var a int
	//fmt.Println("请输入a:")
	//fmt.Scan(&a)
	//fmt.Println("a=",a)

	path := "./demo.txt"
	WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}

func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer f.Close()

	//新建缓冲区,将文件内容放到缓冲区
	r:=bufio.NewReader(f)

	for {
		buf,err:=r.ReadBytes('\n')	//遇到\n结束读取,但\n符也读取出来了
		if err != nil {
			if err==io.EOF{	//文件结束
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Println("buf=",string(buf))
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer f.Close()
	buf := make([]byte, 1024*2) //2k
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {	//文件出错,同时没有到结尾
		fmt.Println(err1)
		return
	}
	fmt.Println("buf=", string(buf[:n]))

}

func WriteFile(path string) {
	//打开文件,新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用完毕,需要关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		f.WriteString(buf)
	}


}
