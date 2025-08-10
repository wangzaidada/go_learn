package main
// 3340. 检查平衡字符串
import (
	"fmt"
)

func isBalanced(num string) bool {
	sum := 0
	signal := -1 // 符号
	for _, val := range num {
		s := int(val - '0') // ASCII 转数字 
		sum += s * signal // 符号 * 数字 实现奇数位为正，偶数位为负
		signal = -signal // 奇偶位符号转换
	}
	return sum == 0

}
func main() {
	fmt.Println("123 is Balance", isBalanced("123"))
	fmt.Println("24123 is Balance", isBalanced("24123"))
}
