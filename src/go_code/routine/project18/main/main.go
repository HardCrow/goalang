package main

import "fmt"

// 类型断言的机制案例
type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var pio Point = Point{1, 2}
	a = pio
	var b Point
	//	b=a
	b = a.(Point) //类型断言  判断a是否指向point类型 如果是就将a赋给b
	fmt.Println(b)
}
