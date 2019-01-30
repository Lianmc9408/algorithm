package mergeSort

func Sort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)>>1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	marge(arr, left, mid, right)
}

func marge(arr []int, left, mid, right int) {
	size := right - left + 1
	p := left
	p2 := mid + 1
	arr2 := make([]int, size)
	i := 0
	for p <= mid && p2 <= right {
		if arr[p] < arr[p2] {
			arr2[i] = arr[p]
			p += 1
		} else {
			arr2[i] = arr[p2]
			p2 += 1
		}
		i ++
	}
	for p <= mid {
		arr2[i] = arr[p]
		i += 1
		p += 1
	}
	for p2 <= right {
		arr2[i] = arr[p2]
		i += 1
		p2 += 1
	}
	for a := range arr2 {
		arr[left+a] = arr2[a]
	}
}
