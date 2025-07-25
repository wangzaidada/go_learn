package main

import (
	"fmt"
)

// 对象（结构体）
type Animal struct {
	name string
}

// 方法
func (a *Animal) Eat() {
	fmt.Println("%s is exting", a.name)
}

// 继承
type Dog struct {
	Animal
	Breed string
}

// 重写
func (d Dog) Eat() {
	fmt.Printf("%s (a %s) is eating dog food\n", d.name, d.Breed) // Printf 支持格式化输出，println 不支持！！
}

// 多态 （接口实现）
type Speaker interface {
	Speak() string
}
type Cat struct {
	Animal
}

// 接口实现多态
func (c Cat) Speak() string {
	return "Miao"
}
func (d Dog) Speak() string {
	return "Wang"
}
func MackSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{
		Animal: Animal{name: "buddy"},
		Breed:  "Golden Retriever",
	}
	dog.Eat()
	MackSound(dog)
	cat := Cat{
		Animal: Animal{name: "kiti"},
	}
	MackSound(cat)
}
