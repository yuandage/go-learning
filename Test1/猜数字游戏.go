package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano()) //设置种子
	var num int
	for {
		num = rand.Intn(10000)
		if num > 1000 {
			break
		}
	}
	*p = num
}

func GetNum(s []int, num int) {
	s[0] = num / 1000       //取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10   //取十位
	s[3] = num % 10         //取个位
}

func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个4位数")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}
		//fmt.Println("num=", num)

		GetNum(keySlice, num)
		//fmt.Println("keySlice=", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 {
			fmt.Println("全部猜对了!!!")
			break
		}
	}
}

func main() {
	var randNum int
	CreateNum(&randNum) //创建一个4位随机数
	//fmt.Println("randNum:", randNum)

	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice=", randSlice)

	OnGame(randSlice)
}
