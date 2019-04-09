package main

import "fmt"

/**
题目：
实现快排

*/
func main() {
	arr := []int{7,2,5,1,3,6,4}
	quickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left int, right int) {
	if right > left {
		pos := partation(arr, left,right)
		quickSort(arr, left, pos-1)
		quickSort(arr, pos+1, right)
	}

	return
}

func partation(arr []int, left int, right int) int {
	key := arr[right]
	o := left-1

	for i:=left;i<right;i++{
		if arr[i] <= key {
			o++
			arr[i],arr[o] = arr[o],arr[i]
		}
	}

	arr[right],arr[o+1] = arr[o+1],arr[right]
	return o+1
}

