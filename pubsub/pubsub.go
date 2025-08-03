// 发布订阅模型
package utils

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者为一个通道
	topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象

type Publisher struct {
	m           sync.RWMutex  //读写锁
	buffer      int           // 通道缓存
	timeout     time.Duration //
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher { // 新建发布者 类似面向对象的 init 初始化对象
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} { // 全部订阅
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} { //订阅主题
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Evict(sub chan interface{}) { // 取消订阅
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Publish(v interface{}) { //发布一个主题
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers { // 循环发送到每一个订阅者
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) Close() { // 关闭发布者
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) sendTopic( // 发布消息
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) { // 过滤不满足条件的
		return
	}
	select {
	case sub <- v: // 发送消息
	case <-time.After(p.timeout): // 超时
	}
}
