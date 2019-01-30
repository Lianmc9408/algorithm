package smallSum

import (
	"algorithm/utils"
	"fmt"
	"testing"
)

func TestSmallSum(t *testing.T) {
	for i := 0; i < 5000; i ++ {
		arr := utils.GenerateArray(20, 20)
		if stupidMethod(arr) != smallSum(arr) {
			fmt.Println("Answer: ", stupidMethod(arr))
			fmt.Println("算法: ", smallSum(arr))
			return
		}
	}
}

func stupidMethod(arr []int) int {
	result := 0
	for i := 1; i < len(arr); i ++ {
		for j := 0; j < i; j ++ {
			if arr[j] < arr[i] {
				result += arr[j]
			}
		}
	}
	return result
}
