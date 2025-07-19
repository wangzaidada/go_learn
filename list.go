package main

import "fmt"

func main() {
	var a [3]int
	var b = [...]int{1, 2, 3} // 直接赋值
	var c = [...]int{2: 3, 1: 2} //索引赋值方式
	var d = [...]int{1, 2, 4: 5, 6} // 混合
	
	// 使用%v占位符打印数组内容
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
}