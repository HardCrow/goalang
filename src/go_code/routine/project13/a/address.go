package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	//结构体赋值方式一
	p2 := person{"huang", 22}
	//结构体赋值方式二
	var p1 *person = new(person)
	p1.Name = "js" //p1.name="xx" 底层就是(*p1).Name="xx"的形式 前提是必须 var p1 *person =new （person）
	//(*p1).Name="li"
	(*p1).Age = 20

	fmt.Println(p1)
	fmt.Println(p2)
}
