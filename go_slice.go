// 切片是简化动态数组。
package main
import (
	"fmt"
)

func TrimSpace(s []byte) []byte{
		b:= s[:0]
		for _, x := range(s){
			if x != ' '{
				b = append(b, x)
			}
		}
		return b
	}

func main(){
	// 切片定义
	var (
	a1 []int	// nil切片,不存在的切片，和nil相等
	b = []int{}	// 空切片，表示一个空集合
	c = []int{1, 2, 3}// len=cap=3
	// d = c[:2]	// len=2,cap=3
	// e = c[0:2:cap(c)]	// len=2; cap=3
	// f = c[:0]	// len=0; cap=3
	// g = make([]int, 3)	// len = cap=3
	// h = make([]int, 2, 3)// len = 2,cap =3
	// i = make([]int, 0, 3)// len =0, cap =3
	)
	// 切片遍历
	fmt.Println("遍历元素")
	for i := range a1 { // 遍历key
		fmt.Printf("a1[%d]: %d\n", i,a1[i])
	}
	for key, value := range b {
		fmt.Printf("b[%d]: %d\n", key, value)
	}
	for i:=0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
	// 添加切片
	fmt.Println("添加元素")
	var a []int
	a = append(a, 1)				//追加一个元素
	a = append(a, 1, 2, 3)			// 追加多个元素，手写解包方式
	a = append(a, []int{1, 2, 3}...) 	// 追加一个切片，切片需要解包#... 表示解包展开切片
	fmt.Println(a) 
	// 开头添加 # 性能较差相比尾插
	a = append([]int{0}, a...)			// 开头添加一个元素			#... 表示解包展开切片
	a = append([]int{-3,-2,-1}, a...)	//开头添加一个切片			#... 表示解包展开切片
	fmt.Println(a)
	// 切片头插会导致内存从新分配，
	// 中间插入(中间操作会产生临时变量)
	var a2 = []int{1,2,3}
	i := 1
	a2 = append(a2[:i], append([]int{999}, a2[i:]...)...) 

	fmt.Println(a2)
	a2 = []int{1,2,3}
	a2 = append(a2[:i], append([]int{9,9,9}, a2[i:]...)...)
	fmt.Println(a2)
	// 中间插入（不产生临时变量）
	a2 = []int{1,2,3}
	a2 = append(a2, 0)
	copy(a2[i+1:], a2[i:]) // 将 i位后面的元素复制到 i+1，相当于i位后移一位
	a2[i] = 999
	fmt.Println(a2)
	// 删除元素
	fmt.Println("删除元素 略")
	// 切片内存技巧
	s := "hello world"
	fmt.Println(string(TrimSpace([]byte(s))))
	// 避免内存溢出
	fmt.Println("垃圾回收")
	o, p, q := 1, 2, 3
	var x = []*int{&o,&p,&q}
	x = x[:len(x)-1] // 删除最后一个元素，未回收
	x = []*int{&o,&p,&q}
	x[len(x) - 1] = nil // 将需要回收的指针设置为nil，包装垃圾回收器可以发现
	x = x[:len(x)-1] // 可回收
	// 强制
}

