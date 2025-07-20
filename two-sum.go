package main
import (
	"fmt"
)
func twoSum(nums []int, target int) []int {
    mp:=map[int]int{}
    for i:=0;i<len(nums);i++ {
        if x,ok :=mp[target-nums[i]]; ok {
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