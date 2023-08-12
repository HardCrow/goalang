package main

import (
	"testing"
)

// 编写一个函数测试main.go里面的函数是否错误
func TestAdd(t *testing.T) { //命名规范和java一样 传参一定要(t *testing.T)
	res := Add(10)
	if res != 55 {
		t.Fatalf("错误")
	} else {
		t.Logf("成功：%v\n", res)
	}
}
