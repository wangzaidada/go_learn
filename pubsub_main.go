package main

import (
	"fmt"
	"strings"
	"time"

	utils "pubsub.go/pubsub"
)

func main() {
	p := utils.NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golong := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golong")
		}
		return false
	})
	p.Publish("hello, world!")
	p.Publish("hello, golong!")
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golong {
			fmt.Println("golong:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
