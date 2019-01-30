package smallSum

import (
	"fmt"
	"testing"
)

func TestHolandFlag(t *testing.T) {
	arr := []int{5, 5, 8, 4, 6, 7, 3, 4, 5, 8, 2, 4, 6, 8, 9, 4, 5, 6, 8, 9, 2}
	division := 8
	holandClassify(arr, division)
	fmt.Println(arr)
}
