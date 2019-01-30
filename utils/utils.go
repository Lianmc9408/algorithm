package utils

import (
	"math/rand"
	"time"
)

func GenerateArray(size, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(max + 1)
	}
	return arr
}

func ArrIsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

