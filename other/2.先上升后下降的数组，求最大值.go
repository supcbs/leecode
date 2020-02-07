package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3, 2, 1, 0}
	maxNum := getMax(arr, 0, len(arr))
	fmt.Println(maxNum)
}

func getMax(arr []int, i int, j int) int {
	mid := (i + j) / 2

	if mid == 0 || mid == len(arr)-1 {
		return arr[mid]
	}

	if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
		return arr[mid]
	}

	if arr[mid] > arr[mid+1] {
		return getMax(arr, i, mid)
	}
	if arr[mid] > arr[mid-1] {
		return getMax(arr, mid, j)
	}
	return mid
}
