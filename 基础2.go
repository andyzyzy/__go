package main

import (
	"fmt"
	"time"
)

func Chan(ch chan int, stopCh chan bool) {
	for i := 1; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func main() {
	ch := make(chan int)
	stopCh := make(chan bool)
	go Chan(ch, stopCh) // 并行
	for {
		select {
		case i := <-ch:
			fmt.Println("接受", i)
		case j := <-ch:
			fmt.Println("接受2", j)
		case _ = <-stopCh:
			goto end // 结束
		}
	}
end:
}
