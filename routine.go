package main

import (
	"fmt"
	"sync"
)

// 原子操作
var total struct { // 声明了一个锁结构体（匿名写法）
	sync.Mutex
	value int
}

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		total.Lock()
		total.value += 1
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Worker(&wg)
	go Worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}
