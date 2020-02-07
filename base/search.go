package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := bSearch(arr, 3)
	fmt.Println(r)
}

// 二分查找
func bSearch(arr []int, num int) int {
	if len(arr) < 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1
	for start <= end {
		// 这一步关键，起始位置+（末尾位置-起始位置）/2
		mid := start + (end-start)/2
		if arr[mid] == num {
			return mid
		} else if arr[mid] > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
