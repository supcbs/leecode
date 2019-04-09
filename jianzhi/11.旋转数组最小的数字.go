package main

import "fmt"

/**
题目：
给一个递增的数组，翻转其中一段，求出翻转的点。
例：[1,2,3,4,5,6] => [4,5,6,1,2,3] => 1

分析：
直接遍历一遍肯定可以。又没有更优的解？

由于考虑到是递增数组，可以考虑使用二分的方式。
index1、index2、mid

当nums[index1] <= index1[mid]时候，说明翻转点在后头
当nums[index1] > index1[mid]时候，说明翻转点在牵头

终止条件：当index1和mid只差一个的时候就终止。
执行条件：当nums[index1] >= nums[index2]的时候就执行。

需要考虑一种情况： 10111和11101
无法通过大小来定位二分缩小的方向，则只能进行遍历了
*/
func main() {
	arr := []int{4, 5, 6,7, 1, 2, 3}
	arr = []int{1,0,1,1,1}
	r := minNumberInRotatedArray(arr)
	fmt.Println(r)
}

func minNumberInRotatedArray(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	index1 := 0
	index2 := len(nums) - 1
	mid := index1

	if nums[index1] >= nums[index2] {
		if index2-index1 == 1 {
			mid = index2
			return mid
		}

		if nums[index1] == nums[index2] && nums[index1] == nums[mid] {
			return findInOrder(nums, index1, index2)
		}

		mid = (index1 + index2) / 2

		if nums[index1] > nums[mid] {
			index2 = mid
		} else {
			index1 = mid
		}
	}

	return nums[mid+1]
}

func findInOrder(nums []int, index1, index2 int) int {
	if len(nums) <= 0 {
		return -1
	}

	ret := nums[index1]
	for i := index1; i <= index2; i++ {
		if ret > nums[i] {
			ret = nums[i]
		}
	}

	return ret
}
