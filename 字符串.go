package main
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	var data = [...]byte{'h','e','l','l','o',',','w','o','r','l','d',}
	fmt.Println("%v",data)
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello)
	fmt.Println(world)
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):",(*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1):",(*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s2):",(*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)
	// go 中字符串为静态，如果要修改，需要先转换为可变的切片类型，修改之后在转回字符串类型。
	// 转换为 []rune 或 []byte
}