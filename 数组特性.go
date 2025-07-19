package main
import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main(){
	var a = [...]int{1,2,3}
	var b = &a
	fmt.Println(a[0], a[1]) // b为指针索引结果一致
	fmt.Println(b[0], b[1])
	for i, v := range b { // key, value 遍历方式
		fmt.Println(i, v)
	}
	for i := 0 ; i < len(b); i++ { // 通过 len()获取数组长度来遍历数组
		fmt.Println(i, b[i])
	}
	var times [5][0]int // 内存占用大小为0 不消耗内存迭代
	for range times {
		fmt.Println("hello")
	}
	// 字符串数组
	var x = [2]string{"hello", "world"}
	var y = [...]string{"你好", "世界"}
	var z = [...]string{1:"世界", 0:"你好"}
	fmt.Println(x, y, z)
	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y : 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1, line2, line3)
	// 函数数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Println(decoder1, decoder2)]
	// 接口数组
	var unkonwn1 [2]interface{}
	var unkonwn2 = [...]interface{}{123, "你好"}
	fmt.Println(unkonwn1, unkonwn2)
	// 通道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %V\n", b)
}