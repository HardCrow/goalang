package main

import (
	"testing"
)

// 编写一个函数测试main.go里面的函数是否错误
func TestSub(t *testing.T) { //命名规范和java一样 传参一定要(t *testing.T)
	res := getSub(3, 10)
	if res != 7 {
		t.Fatalf("错误")
	} else {
		t.Logf("成功")
	}
}
