package main

import "fmt"

type A struct {
	Name string
	age  int
}
type B struct {
	Name  string
	score float64
}
type C struct {
	A
	B
	// Name string
}
type D struct {
	a A //有名结构体   组合关系
}

func main() {
	var c C
	//	c.Name="aaa"   //错误 ab都有name属性  c结构体有自己的name属性就不会出错
	//
	// 或者下面方法指定那个结构体的name
	c.B.Name = "aaaa"
	c.A.Name = "bbb"
	var d D
	// d.Name ="ada"  //d不是匿名结构体 因此不会在a结构体里面找 解决方式如下   如果结构体d里面有name就用d结构体里面的属性
	d.a.Name = "sd"
	fmt.Println()
}
