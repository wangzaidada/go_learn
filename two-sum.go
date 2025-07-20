package main
import (
	"fmt"
)
func twoSum(nums []int, target int) []int {
    mp:=map[int]int{}
    for i:=0;i<len(nums);i++ { // 这里的len() 返回已使用的元素数，cap()返回最大可容纳的元素数。注意map是动态类型,不适用cap
        if x,ok :=mp[target-nums[i]]; ok { // 如果x + y = z的话 y = z - x成立；
            return []int{x,i}
        }
        mp[nums[i]]=i
    }
    return []int{}
}
func main(){
	var nums = []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums,target))
}