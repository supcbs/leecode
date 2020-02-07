package main

import "fmt"

/**

题目:
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定 nums = [1,1,1,2,2,3],
函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,1,2,3,3],
函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。
你不需要考虑数组中超出新长度后面的元素。

说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

func main() {
	nums := []int{1,1,1,2,2,3}
	r := removeDuplicatesii(nums)
	fmt.Println(r)
}

/**
方法1:
这边和26题很像，不过是这边允许连续两位一样。
注意题目是排序数组，说明是递增的。

参考26题的方法，采用双指针。
这边只需要满足j>i-2即可说明是一个新的数字了，然后赋值给i，i递增1位。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeDuplicatesii(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 2 {
		return 2
	}

	i := 0
	for j := 0; j < numsLen; j++ {
		if i < 2 || nums[j] > nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}


/**
举一反三：

这道题核心因素还是顺序排列。不管允许出现n次。用双指针法进行遍历的时候，当大于n的时候，偏移量j一定大约j-n，否则进行替换。

 */
