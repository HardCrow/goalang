package main

import "fmt"

func main() {
	//string字符串的修改
	//原因：string字符串在底层是一个数组并且string不可变的不能通过str[0]=‘z’的方式去改变
	str := "huang"
	slice := str[1:]
	fmt.Println(slice) //证明string类型的字符串底层是数组类型
	arr := []byte(str) //转成byte切片  如果想要修改成汉字变成[]rune(str)即可也就是类型改成rune类
	arr[0] = 'A'       //切片修改
	str = string(arr)  //数组转化
	fmt.Println(str)
}
