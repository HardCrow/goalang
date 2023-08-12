package main

import (
	"fmt"
	"go_code/routine/project21-test/work"
	"testing"
)

// 测试类必须 xxx_test命名
func TestStore(t *testing.T) {
	monster := work.Monster{
		Name:  "黄",
		Age:   99,
		Skill: "s",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("错误%v 实际%v", true, res)
	}
	t.Log("成功")
	resr := monster.Restore()
	fmt.Println(resr)
}
