package main

import "fmt"

func main() {
	n := fbnq(19)
	fmt.Println(n)

	arr := [6]int{1, 6, 8, 52, 452}
	BirnaryFind(&arr, 0, len(arr)-1, 8)
	BirnaryFind(&arr, 0, len(arr)-1, 452)
	BirnaryFind(&arr, 0, len(arr)-1, 0)
}
func BirnaryFind(arr *[6]int, leftIndex int, rightIndex int, finalVar int) { //二分法
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	middle := (rightIndex + leftIndex) / 2
	if (*arr)[middle] > finalVar {
		BirnaryFind(arr, leftIndex, middle-1, finalVar)
	} else if (*arr)[middle] < finalVar {
		BirnaryFind(arr, middle+1, rightIndex, finalVar)
	} else {
		fmt.Printf("下标为%v \n", middle)
	} //二分法
}
func fbnq(n int) []uint64 {
	fbnqSlice := make([]uint64, n) //声明一个切片
	//第一个数和第二个数都是已经写好的
	fbnqSlice[0] = 1
	fbnqSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnqSlice[i] = fbnqSlice[i-1] + fbnqSlice[i-2]
	}
	return fbnqSlice
}
