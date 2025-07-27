// 通道，线程通信，实现线程同步
package main

import (
	. "fmt"
)

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好，世界"
	done <- true
	close(done)
}

func main() {
	go aGoroutine()
	<-done
	Println(msg)
}
