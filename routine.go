package main

import (
	"fmt"
	"sync"
)

// 原子操作
var total struct { // 声明了一个锁结构体（匿名写法）
	sync.Mutex //加锁确保资源访问互斥
	value      int
}

func Worker(wg *sync.WaitGroup) {
	defer wg.Done() // 确保函数退出时通知等待组
	for i := 0; i < 100; i++ {
		total.Lock() // 资源访问互斥
		total.value += 1
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup // 等待组
	wg.Add(2)             // 注册两个任务
	go Worker(&wg)        // 启动任务，传入等待组指针
	go Worker(&wg)
	wg.Wait() // 阻塞等待所有注册的goroutine完成
	fmt.Println(total.value)
}
