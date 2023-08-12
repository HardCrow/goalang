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
		Skill: "吃饭",
	}
	res := monster.Store()
	fmt.Println(res)
	resr := monster.Restore()
	fmt.Println(resr)
}
