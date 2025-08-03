package main

import (
	"fmt"
	"strings"
	"time"

	pubsub "pubsub.go/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	all := p.Subscribe()                                  // 订阅全部内容
	golong := p.SubscribeTopic(func(v interface{}) bool { //传入消息过滤判断器
		if s, ok := v.(string); ok { // 如果消息类型是字符型
			return strings.Contains(s, "golong") // 订阅 golong true, 其他 false
		}
		return false // 不订阅  false
	})
	p.Publish("hello, world!")
	p.Publish("hello, golong!")
	go func() { // 启动 匿名函数
		for msg := range all { // 持续从通道中获取数据
			fmt.Println("all:", msg) // 输出数据
		}
	}()

	go func() {
		for msg := range golong { // 持续从通道中获取数据
			fmt.Println("golong:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
