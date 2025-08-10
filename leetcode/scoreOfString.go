package main
//3110. 字符串的分数
import "fmt"

// scoreOfString 计算字符串的分数
// 分数定义为相邻字符ASCII值差值的绝对值之和
// 参数: s - 输入的字符串
// 返回: 计算得到的分数
func scoreOfString(s string) int {
	// 获取字符串长度
	num := len(s)
	// 初始化分数为0
	score := 0
	// 遍历字符串，比较相邻字符
	for i := 0; i+1 < num; i++ {
		// 计算当前字符与下一个字符的ASCII值差值
		score_str := int(s[i]) - int(s[i+1])
		// 如果差值为负数，取绝对值
		if score_str < 0 {
			score_str = -score_str
		}
		// 将差值累加到总分中
		score += score_str
	}
	return score
}

// main 主函数，演示scoreOfString函数的使用
func main() {
	// 计算字符串"hello"的分数并打印结果
	fmt.Printf("%d", scoreOfString("hello"))
}
