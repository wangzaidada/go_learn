package main

// 单例模式实现方式
import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32     // 是否初始化标记
	mu          sync.Mutex // 加锁
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 { // 检查对象是否已经实例化
		return instance // 已经实例化直接返回
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil { // 双重检查
		defer atomic.StoreUint32(&initialized, 1) // 原子操作标记已实例化
		instance = &singleton{}                   // 实例化
	}
	return instance
}

// 单例实现标准库
// 基于sync.Once 的单例实现
var (
	instance_v2 *singleton
	once        sync.Once
	initcount   uint64
)

func Instance_v2() *singleton {
	once.Do(func() {
		instance_v2 = &singleton{}
		atomic.AddUint64(&initcount, 1)
	})
	return instance_v2
}

func main() {
	fmt.Print(Instance())
	x := 10
	var mu sync.WaitGroup
	mu.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			defer mu.Done()
			inst := Instance_v2()
			fmt.Printf("实例地址:%p\n", inst)
		}()
	}
	mu.Wait() //如果忘记Wait 会导致程序提前结束
	fmt.Print(initcount)
}
