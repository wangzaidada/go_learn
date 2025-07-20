package main
import "fmt"
func isPalindrome(x int) bool {
    if x <0 || (x % 10 ==0 && x != 0){ // 负数必不回文，非零的零结尾必不回文
        return false
    }

    revertedNumber := 0
    for x > revertedNumber {
        revertedNumber = x % 10 + revertedNumber * 10
        x /= 10
    }
    return x == revertedNumber || x == revertedNumber / 10 // 偶数位会直接相等，奇数位在/10 之后也会相等
}
func main(){
	x :=  12321
	fmt.Println(isPalindrome(x))
}