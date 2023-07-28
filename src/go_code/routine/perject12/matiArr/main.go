package main

import "fmt"

func main() {
	//定义二维数组
	var arr [4][6]int
	fmt.Println(arr)
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 9, 6}} //自己定义数组值操作
	fmt.Println(arr2)
}
