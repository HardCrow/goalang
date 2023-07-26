package main

import "fmt"

func main() {
	// 切片的基本使用
	var intArr [5]int = [5]int{2, 62, 85, 3, 40}
	//声明切片
	//slice 切片名
	//intArr【1：3】表示slice引用到intArr这个数组
	//引用intArr数组的起始下标为1最后的下标为三 不包括三
	slice := intArr[1:3]
	fmt.Println(intArr)
	fmt.Println(slice)      //内容
	fmt.Println(len(slice)) //元素个数
	fmt.Println(cap(slice)) //容量
}
