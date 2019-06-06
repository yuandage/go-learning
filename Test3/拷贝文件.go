package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) == 1 {
		fmt.Println("请在终端运行! 或输入命令参数!")
		fmt.Println("程序将在3s内退出!")
		time.Sleep(3*time.Second)
		return
	}
	if len(list) != 3 {
		fmt.Println("example: copy.exe srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}
	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}

	//操作完毕,关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理,从源文件读取内容,写入目的文件
	buf := make([]byte, 1024*4) //4k大小缓冲区
	for {
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF { //文件读取完毕
				fmt.Println("文件复制成功!")
				break
			}
			fmt.Println("err=", err)
		}
		//写入目的文件
		dF.Write(buf[:n])
	}
}
