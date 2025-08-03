package main

import "fmt"

func GenerateNatural() chan int { // 生产从2开始的自然数
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}

		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 通道第一个数必为素数，依次构建不能被2整除的通道，第二次得到的是不能被2和3 整除的通道依次类推
	}
}
