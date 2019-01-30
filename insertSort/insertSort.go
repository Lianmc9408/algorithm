package insertSort

// 把最后一个数插入到前面
func Sort(arr []int) {
	for i := 1; i < len(arr); i ++ {
		for j := i - 1; j >= 0; j -- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				break
			}
		}
	}
}
