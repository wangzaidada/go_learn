package main

// 2614. 对角线上的质数
import "fmt"

func is_Prime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func diagonalPrime(nums [][]int) int {
	sum := len(nums)
	maxPrime := 0
	for i := 0; i < sum; i++ {
		val1 := nums[i][i]
		if is_Prime(val1) && val1 > maxPrime {
			maxPrime = val1
		}

		val2 := nums[i][sum-1-i]
		if i != sum-1-i && is_Prime(val2) && val2 > maxPrime {
			maxPrime = val2
		}
	}
	return maxPrime
}
func main() {
	matrix := [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}
	fmt.Printf("%d", diagonalPrime(matrix))
}
