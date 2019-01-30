package main

import (
	"algorithm/mergeSort"
	"algorithm/utils"
	"fmt"
)

func main() {
	arr := utils.GenerateArray(8, 50)
	fmt.Println(arr)
	//bubbleSort.Sort(arr)
	//insertSort.Sort(arr)
	//chooseSort.Sort(arr)
	//quickSort.Sort(arr)
	mergeSort.Sort(arr)
	fmt.Println(arr)
}