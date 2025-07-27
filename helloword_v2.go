package main

// 锁实现线程同步
import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock() // 加锁保证程序不会提前结束
	go func() {
		fmt.Println("hello world")
		mu.Unlock() // 解锁，保证程序正常结束
	}()
	mu.Lock() // 在已有锁的情况下会等待
}
