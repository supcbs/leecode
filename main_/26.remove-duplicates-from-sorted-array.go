package main

import "fmt"

/**
题目:
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。

说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r := removeDuplicates(nums)
	fmt.Println(r)
}

/**
方法1:
双指针法。
通过ij指针，i=0,j=1分别向前移动。
当i=j的时候，说明重复，则不作处理
当i!=j的时候，将i复制给j+1位

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}

	var i,j int
	for ; i < numsLen; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
