package smallSum

// 快速排序类型: 荷兰国旗问题, 也称三色排序问题
// 给定一个数组和一个数字, 把小于这个数的放数组左边, 大于这个数的放数组右边, 等于的放中间

func holandClassify(arr []int, division int) {
	left := 0
	right := len(arr) - 1
	index := 0
	for index <= right {
		if arr[index] > division {
			arr[index], arr[right] = arr[right], arr[index]
			right --
		} else if arr[index] == division {
			index += 1
		} else {
			if left != index {
				arr[left], arr[index] = arr[index], arr[left]
			}
			left ++
			index ++
		}
	}
}
