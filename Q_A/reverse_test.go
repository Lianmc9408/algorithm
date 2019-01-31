package Q_A

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) int {
	var r int
	for x != 0 {
		n := x % 10
		x = x / 10
		fmt.Println(n)
		if r > math.MaxInt32/10 || (r == math.MaxInt32/10 && n > 7) {
			return 0
		}
		if r < math.MinInt32/10 || (r == math.MinInt32/10 && n < -8) {
			return 0
		}
		r = r*10 + n
	}
	return r
}

func romanToInt(s string) int {
	dicts := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	last, r := 0, 0

	for _, n := range s {
		num := dicts[string(n)]
		if num > last {
			last = num - last
		} else {
			r += last
			last = num
		}
	}
	r += last
	return r
}

func isValid(s string) bool {
	arr := make([]string, len(s)/2)
	symbolDict := map[string]int{
		"(": 1,
		"{": 2,
		"[": 3,
		")": -1,
		"}": -2,
		"]": -3,
	}
	index := 0
	for _, symbol := range s {
		sy := string(symbol)
		if symbolDict[sy] > 0 {
			arr[index] = sy
			index ++
		} else {
			if index <= 0 {
				return false
			}
			if symbolDict[sy]+symbolDict[arr[index-1]] != 0 {
				return false
			}
			index --
		}
	}
	return index == 0
}

func isValidV2(s string) bool {
	arr := make([]string, 0)
	symbolDict := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for _, symbol := range s {
		sy := string(symbol)
		if ss, ok := symbolDict[sy]; !ok {
			arr = append(arr, sy)
		} else if len(arr) == 0 {
			return false
		} else if ss != arr[len(arr)-1] {
			return false
		} else {
			arr = arr[0 : len(arr)-1]
		}
	}
	return len(arr) == 0
}

func TestReverse(t *testing.T) {
	//fmt.Println(reverse(1534236469))
	//fmt.Println(romanToInt("IV"))
	fmt.Println(isValidV2("[[{}]]"))

	// ()()()()(){}][
}
