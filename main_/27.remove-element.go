package main

import "fmt"

/**
题目:

给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。
*/

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	r := removeElement(nums, 2)
	fmt.Println(r)
}

/**
如果非val的元素在最后，那么必须扫描到最后才知道，对之前的元素做了不必要的替换操作。
如果非val的元素在第一个，那么也似乎没必要将后面的都左移赋值。

所以来了一种思路：
将当前扫描到的元素若==val，则和最后一个元素替换，同时缩小数组长度；
若当前扫描到的元素若!=val，则指针右移。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeElement(nums []int, val int) int {
	var i int
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}

	for i < numsLen {
		if nums[i] == val {
			// 这里是减1，因为数组的商都一直在缩减
			nums[i] = nums[numsLen-1]
			numsLen--
		} else {
			i++
		}
	}

	return i
}
