package main

import (
	"fmt"
	"sync"
	"time"
)

func woker(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		select { // select 是 “通道操作就绪检测器”，switch 是 “值的条件判断器”。
		case <-cancel:
			fmt.Println("关闭...")
			return
		default: // 条件不满足执行
			fmt.Println("运行中...")
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go woker(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()

}
