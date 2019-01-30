package smallSum

// 归并排序类型
// 小和问题, 在一个数组中,每一个元素左边比当前元素值小的元素值累加起来,叫做这个数组的小和
func smallSum(arr []int) int {
	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) int {
	if left == right {
		return 0
	}
	mid := left + (right-left)>>1
	return mergeSort(arr, left, mid) + mergeSort(arr, mid+1, right) + merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) int {
	result := 0
	size := right - left + 1
	arr2 := make([]int, size)
	p := left
	p2 := mid + 1
	index := 0
	for p <= mid && p2 <= right {
		if arr[p] < arr[p2] {
			result += arr[p] * (right - p2 + 1)
			arr2[index] = arr[p]
			p += 1
		} else {
			arr2[index] = arr[p2]
			p2 += 1
		}
		index += 1
	}
	for p <= mid {
		arr2[index] = arr[p]
		p += 1
		index += 1
	}
	for p2 <= right {
		arr2[index] = arr[p2]
		p2 += 1
		index += 1
	}
	for i := range arr2 {
		arr[left+i] = arr2[i]
	}
	return result
}
