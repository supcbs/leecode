package main

import "fmt"

/**

题目:
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的原地算法。
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	//rotate1(nums, k)
	//rotateOK(nums, k)
	//rotate3(nums, k)
	rotateNew(nums, k)
	fmt.Println(nums)
}

/**
go当中可以直接拼接字符串，然后赋值即可
 */
func rotateNew(nums []int, k int) {
	numsLen := len(nums)
	if numsLen <= 0{
		return
	}
	numK := k % numsLen
	a := append(nums[numsLen-numK:],nums[:numsLen-numK]...)

	for i:=0;i<len(nums);i++{
		nums[i] = a[i]
	}
}



/**
方法1:
最能想到的就是每次右移1位，然后移动k次。
每次右移1位：将最后一位(len-1)存住temp，然后从大到小将j赋值给j-1，最后temp赋值第0位即可

关键点：倒着赋值，正向赋值会被覆盖

时间复杂度：O(kn)
空间复杂度：O(1)
*/
func rotate1(nums []int, k int) {
	numsLen := len(nums)
	k %= numsLen

	for i := 0; i < k; i++ {
		temp := nums[numsLen-1]
		for j := numsLen - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = temp
	}
}

/**
方法2:
翻转的思路，假如是12345
a.现将全部翻转54321
b.将前k位翻转34521
c.将k位后翻转34512
即可实现右移的效果

时间复杂度：O(n)
空间复杂度：O(1)
*/
func rotateOK(nums []int, k int) {
	numsLen := len(nums)
	k %= numsLen

	rotate2Reverse(nums, 0, numsLen-1)
	rotate2Reverse(nums, 0, k-1)
	rotate2Reverse(nums, k, numsLen-1)
}

/**
实现翻转，思路很简单，采用双指针将最后一位和第一位进行交换
*/
func rotate2Reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/**
方法3:
k位循环交换的思路，假如是12 3 45，k=2
a.45 3 12
b.45 132 =>45123
即可实现右移的效果

这里需要关注的是边界调节和终止条件

时间复杂度：O(n^2/k)
空间复杂度：O(1)
*/
func rotate3(nums []int, k int) {
	numsLen := len(nums)
	numsLenC := numsLen
	k %= numsLen

	i := 0
	for i < numsLen && k != 0 {
		for j := 0; j < k; j++ {
			nums[i+j], nums[numsLen-k+j] = nums[numsLen-k+j], nums[i+j]
		}

		numsLenC -= k
		i += k
		k %= numsLenC
	}
}
