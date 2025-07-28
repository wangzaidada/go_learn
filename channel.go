// 通道，线程通信，实现线程同步
package main

import (
	"fmt"
	. "fmt"
)

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好，世界"
	done <- true
	close(done)
}
func ch() { // 多线程同步（通道方式）
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Printf("hello.world %d", i)
			done <- 1
		}()
	}
	for i := 0; i < cap(done); i++ {
		<-done
		fmt.Println("ok")
	}
}

func main() {
	go aGoroutine()
	<-done
	Println(msg)
	ch()
}
