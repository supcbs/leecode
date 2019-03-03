package main

import "fmt"

/**

题目:
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

*/

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroesOK(nums)
	fmt.Println(nums)
}

/**
方法1:

两个移动的指针：
其中一个遇到0则停住，另一个继续往前，遇到非0的则与第一个指针指的数替换，同时第一个指针++
当遇到0

时间复杂度：O(n)
空间复杂度：O(1)
*/
func moveZeroes(nums []int) {
	numsLen := len(nums)

	if numsLen <= 0 {
		return
	}

	var i, j int


	for j < numsLen {

		// 找到指向0的位置
		for nums[i] != 0 && i < numsLen && i != -1 {
			i++
			if i >= numsLen {
				goto LOOP
			}
		}

		// 找到非0的位置，寻找其之前的为0的位置
		if nums[j] != 0 && j > i {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			if i >= numsLen {
				goto LOOP
			}
		}
		j++
	}
LOOP:
	return
}
/**
方法2:

看到一个很秀的写法：
将非0的移到前头，剩下的都置为0

时间复杂度：O(n+?)
空间复杂度：O(1)
*/
func moveZeroesOK(nums []int) {
	numsLen := len(nums)

	if numsLen <= 0 {
		return
	}

	var i,j int


	for j < numsLen{
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	for i < numsLen{
		nums[i] = 0
		i++
	}
	return
}
