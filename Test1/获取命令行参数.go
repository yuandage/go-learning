package main
//导入包方式 比较常用
import (
	"fmt"
	"os"
)

//import io "fmt"	//给包名起别名
//import _ "fmt"	//忽略包 调用init()

//import "fmt"
//import "os"

func main() {
	list:=os.Args
	n:=len(list)

	fmt.Println("n=",n)
	for i,data:=range list {
		fmt.Printf("list[%d]=%s\n",i,data)
	}

}
