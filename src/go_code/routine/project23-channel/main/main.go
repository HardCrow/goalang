package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	aMap = make(map[int]int, 10)
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock() //加锁 上锁
	aMap[n] = res
	lock.Unlock() //解锁
}
func main() {
	for i := 0; i <= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	for i, v := range aMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}

}
