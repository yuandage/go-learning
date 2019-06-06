package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(2) //指定以2核运算
	fmt.Println("n=", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
