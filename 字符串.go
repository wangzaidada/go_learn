package main
import (
	"fmt"
)

func main(){
	var data = [...]byte{'h','e','l','l','o',',','w','o','r','l','d',}
	fmt.Printf("%v",data)
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello)
	fmt.Println(world)
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):",(*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s):",(*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
}