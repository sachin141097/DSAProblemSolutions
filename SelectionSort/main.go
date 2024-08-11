package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minimum := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minimum] {
				minimum = j
			}
		}

		//swap elements
		arr[i], arr[minimum] = arr[minimum], arr[i]
	}

}
func main() {
	arr := []int{7, 2, 9, 6, 4}
	fmt.Printf("Value of array before sorting is %v\n", arr)
	selectionSort(arr)
	fmt.Printf("Value of array after sorting is %v\n", arr)
}
