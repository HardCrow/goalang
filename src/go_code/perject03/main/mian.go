package main

import "fmt"

var n5 = 100
var n6 = 10
var (
	n7 = 300
	n8 = 400
) //全局变量
func main() {
	var n1, n2 int //一次声明多个变量
	fmt.Println(n1, n2)

	n4, name, n3 := 100, "tom", 886 //方式二
	fmt.Println(n4, name, n3)

	fmt.Println(n5, n6, n7, n8)
	//下面表示指针
	var i int = 110
	var ptr *int = &i

	fmt.Printf("ptr=%v\n", ptr)  //指向的是值的地址
	fmt.Printf("ptr=%v\n", *ptr) //指向的是地址对应的值
	fmt.Printf("ptr=%v\n", &ptr) //取出存ptr地址的地址
}
