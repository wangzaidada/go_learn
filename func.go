package main

import "fmt"

// 具名函数
func Add(a, b int) int {
	return a + b
}

// 匿名函数
var add = func(a, b int) int {
	return a + b
}

// 多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, value := range more {
		a += value
	}
	return a
}

// 给函数的返回值命名
func Find(m map[int]int, key int) (value int, ok bool) { // 安全的查找一个map的键
	value, ok = m[key]
	return // 隐式返回命名返回值 value 和 ok
}

// `defer` 是 Go 语言中一个非常有用的关键字，它用于延迟（defer）函数的执行。
// 被 `defer` 修饰的函数调用会被推迟到包含它的函数返回之前执行。
// 无论包含函数是正常返回，还是由于 panic 导致的异常返回，被延迟的函数都会被执行。
// 这通常用于资源清理、解锁、关闭文件等操作，确保在函数结束时执行必要的清理工作。
// 使用 defer 在return之后修改已命名的返回值
func Inc() (v int) {
	defer func() { v++ }() // () 表示对匿名函数的调用
	return 42
}
func f(x int) *int {
	return &x

}
func g() int {
	var x *int // 书上缺少这一步
	x = new(int)
	return *x
}
func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(add(2, 3))
	fmt.Println(Swap(1, 2))
	fmt.Println(Sum(1, []int{1, 2, 3}...))
	var a = []interface{}{123, "abc"} // 空接口类型， 解包会导致结果不同
	fmt.Println(a...)
	fmt.Println(a)
	fmt.Println(Find(map[int]int{1: 1212, 2: 12122}, 1))
	fmt.Println(Find(map[int]int{1: 1212, 2: 12122}, 3))
	fmt.Println(f(1))
	fmt.Print(g())
}
