package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

// sync/atomic 实现原子操作
var total_2 uint64

func Worker_2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64 // 将i 定义为无符号int64类型
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&total_2, 1) // 更轻量级的互斥访问
	}
}
func main() {
	var wg sync.WaitGroup // 等待组
	wg.Add(2)             // 注册两个任务
	go Worker(&wg)        // 启动任务，传入等待组指针
	go Worker(&wg)
	wg.Wait() // 阻塞等待所有注册的goroutine完成
	var wg_2 sync.WaitGroup
	wg_2.Add(2)
	go Worker_2(&wg_2)
	go Worker_2(&wg_2)
	wg_2.Wait()
	fmt.Println(total.value)
	fmt.Println(total_2)
}
