package chooseSort

// 从后面选一个最小的放进去
func Sort(arr []int) {
	for i := 0; i < len(arr); i ++ {
		min := i
		for j := i; j < len(arr); j ++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}
