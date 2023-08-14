package main

import "fmt"

func write(intchan chan int) {
	for i := 0; i < 50; i++ {
		intchan <- i
		fmt.Println(i)
	}
	close(intchan)
}
func read(intchan chan int, boolchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	boolchan <- true
	close(boolchan)
}
func main() {
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)
	go write(intchan)
	go read(intchan, exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}

}
