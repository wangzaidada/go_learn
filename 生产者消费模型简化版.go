package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 定义配置结构体
type Config struct {
	ServerName string
	MaxRetries int
	Timeout    time.Duration
	Version    int // 版本号，用于跟踪配置更新
}

// 模拟从外部加载配置
func loadConfig() *Config {
	// 实际应用中，这里可能从文件、数据库或远程服务加载配置
	// 为了演示，生成一些随机配置值
	rand.Seed(time.Now().UnixNano())// 随机种子
	return &Config{
		ServerName: fmt.Sprintf("server-%d", rand.Intn(100)),
		MaxRetries: rand.Intn(5) + 1,
		Timeout:    time.Duration(rand.Intn(3)+1) * time.Second,
		Version:    int(time.Now().Unix() % 1000), // 简单的版本号生成
	}
}

// 模拟请求来源
func requests() <-chan int {
	ch := make(chan int) // 创建一个 整型通道
	go func() {
		defer close(ch) // 延迟关闭通道
		// 持续生成请求
		for i := 0; i < 50; i++ {// 生产50 个请求
			ch <- i                            // 发送请求ID
			time.Sleep(200 * time.Millisecond) // 控制请求频率
		}
	}()
	return ch
}

var config atomic.Value

func main() {
	// 初始加载配置
	initialConfig := loadConfig()
	config.Store(initialConfig)
	fmt.Printf("初始配置加载完成，版本: %d\n", initialConfig.Version)

	// 启动配置更新goroutine
	go func() {
		for {
			time.Sleep(3 * time.Second) // 每3秒更新一次配置
			newConfig := loadConfig()
			config.Store(newConfig)
			fmt.Printf("配置已更新，新版本: %d\n", newConfig.Version)
		}
	}()

	// 启动10个请求处理goroutine
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(worke rID int) {
			defer wg.Done()
			for reqID := range requests() {
				// 加载当前配置
				c := config.Load().(*Config)

				// 处理请求（示例逻辑）
				fmt.Printf("工作线程 %d 处理请求 %d，使用配置版本 %d，服务器名: %s，超时: %v\n",
					workerID, reqID, c.Version, c.ServerName, c.Timeout)

				// 模拟处理时间
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// 等待所有工作线程完成
	wg.Wait()
	fmt.Println("所有请求处理完成")
}
