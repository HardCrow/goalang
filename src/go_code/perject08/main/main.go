package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	//数组定义 var 定义名 []类型 = []类型{  }
	var a [2]int //定义
	fmt.Println(a)
	var intArr [5]int = [...]int{1, 5, 2, 9, 41} //定义加赋值
	max := intArr[0]
	for i := 0; i < len(intArr); i++ {
		if max < intArr[i] {
			max = intArr[i]
		}
		fmt.Println(max)
	}
	fmt.Println(".....................")
	fmt.Println(max)
	//for-range 语法
	//for index, value := range array01(){.....}
	var sum [4]int = [...]int{6, 3, 4, 2}
	s := 0
	for _, val := range sum {
		s += val
	}
	fmt.Printf("sum=%v average=%v", s, s/len(sum))

	//var array3 [5]int
	//for i:=0;i<len(array3);i++{
	//	array3[i]=rand.Intn(100)
	//}
	//for i:=len(array3);i<len(array3);i-- {
	//	fmt.Println(array3[i])
	//}
}
