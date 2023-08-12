package main

//
import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 0; i <= 5; i++ {
		fmt.Println("ohhhh" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("num", num)
	test()
	go test() //开启了一个协程
	for i := 0; i <= 5; i++ {
		fmt.Println("ohhhh2222" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
