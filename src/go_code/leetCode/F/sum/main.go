package main

import "fmt"

// 题目要求：
// 给定一个整数nums和一个整数目标值target，在该数组中找出和为目标值target的拿两个整数，并返回他们的下标
// 例如: nums =[2,7,11,15] target=9
// 输出[0,1]
// 因为num[0]+num[1]==9 返回[0,1]
func main() { //打卡第一天7.29 和标准答案的相似度是90%
	var nums [4]int = [4]int{2, 7, 11, 15}
	var target int = 9
	fmt.Println(sum1(nums, target))
}
func sum1(arr [4]int, target int) []int { //注意func 函数名 （形参类型） 返回值类型
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j} //return的时候不知道传什么
			}
		}
	}
	return nil //注意没有找到的话还是要返回一个空值
}
