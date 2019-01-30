package utils

import (
	"algorithm/mergeSort"
	"fmt"
	"sort"
	"testing"
)

func TestAlgorithm(t *testing.T) {
	size := 100
	maxValue := 200
	count := 100000
	for i := 0; i < count; i ++ {
		arr1 := GenerateArray(size, maxValue)
		arr2 := make([]int, size)
		arr3 := make([]int, size)
		copy(arr2, arr1)
		copy(arr3, arr1)
		sort.Ints(arr2)
		// =====中间填要测试的排序算法对arr3排序=====
		//bubbleSort.Sort(arr3)
		//insertSort.Sort(arr3)
		//chooseSort.Sort(arr3)
		//quickSort.Sort(arr3)
		mergeSort.Sort(arr3)
		// =====中间填要测试的排序算法对arr3排序=====
		if !ArrIsEqual(arr2, arr3) {
			fmt.Println("===============================")
			fmt.Println("============算法错误============")
			fmt.Println("=arr1:", arr1, "=")
			fmt.Println("=arr2:", arr2, "=")
			fmt.Println("=arr3:", arr3, "=")
			fmt.Println("===============================")
			return
		}
	}
	fmt.Println("排序算法正确")
}
