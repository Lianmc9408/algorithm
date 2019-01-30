package quickSort

// 选个某个数, 把大于这个数的放右边, 小于这个数的放左边, 等于这个数的放中间
// 荷兰国旗问题, 三色排序
func Sort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r := partition(arr, left, right)
	quickSort(arr, left, l-1)
	quickSort(arr, r+1, right)
}

func partition(arr []int, left int, right int) (int, int) {
	flag := arr[right]
	i := left
	for i <= right {
		if arr[i] < flag {
			if i != left {
				arr[i], arr[left] = arr[left], arr[i]
			}
			left += 1
			i += 1
		} else if arr[i] > flag {
			arr[i], arr[right] = arr[right], arr[i]
			right -= 1
		} else {
			i += 1
		}
	}
	return left, right
}
