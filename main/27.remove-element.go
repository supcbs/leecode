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

说明:

为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//r := removeElement(nums, 2)
	r := removeElementOK(nums, 2)
	fmt.Println(r)
}

/**
方法1:
很容易想到的是，双指针法。
i代表非val的位数
j代表一直向前扫描的位数

ij需要互换来保证前n个数为不等于val的（这里j的速度>=i）

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeElement(nums []int, val int) int {
	var i, j int
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}
	for ; j < numsLen; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

/**
方法2：
如果非val的元素在最后，那么必须扫描到最后才知道，对之前的元素做了不必要的替换操作。
如果非val的元素在第一个，那么也似乎没必要将后面的都左移赋值。

所以来了一种思路：
将当前扫描到的元素若==val，则和最后一个元素替换，同时缩小数组长度；
若当前扫描到的元素若!=val，则指针右移。
*/
func removeElementOK(nums []int, val int) int {
	var i int
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}

	for i < numsLen {
		if nums[i] == val {
			nums[i] = nums[numsLen-1]
			numsLen--
		} else {
			i++
		}
	}

	return i
}
