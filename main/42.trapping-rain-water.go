package main

import "fmt"

/**

题目:
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，
计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，
在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/

func main() {
	nums := []int{2, 0, 2}
	r := trap2(nums)
	fmt.Println(r)
}

/**
方法1:
可以这么理解，
从左往右遍历一遍，取最大值进行后续的赋值，可以理解为左边开始的灌水:010202 -> 011222
从右往左遍历一遍，取最大值进行后续的赋值，可以理解为右边开始的灌水:010202 -> 222222

灌水后，遍历一回数组，取水位低的数，来与当前的值做差值，就是接到的水

时间复杂度：O(n)
空间复杂度：O(1)
*/
func trap(height []int) int {
	heightLen := len(height)
	if heightLen <= 2 {
		return 0
	}

	left := make([]int, heightLen)
	right := make([]int, heightLen)

	// 从左往右遍历一遍，取最大值进行后续的赋值，可以理解为左边开始的灌水:010202 -> 011222
	for i := 1; i < heightLen; i++ {
		if left[i-1] > height[i-1] {
			left[i] = left[i-1]
		} else {
			left[i] = height[i-1]
		}
	}

	// 从右往左遍历一遍，取最大值进行后续的赋值，可以理解为右边开始的灌水:010202 -> 222222
	for j := heightLen - 2; j >= 0; j-- {
		if right[j+1] > height[j+1] {
			right[j] = right[j+1]
		} else {
			right[j] = height[j+1]
		}
	}

	fmt.Println(right, left)
	// 头尾肯定是0不用算
	total := 0
	for i := 1; i < heightLen-1; i++ {
		minY := 0
		if left[i] > right[i] {
			minY = right[i]
		} else {
			minY = left[i]
		}

		if minY > height[i] {
			total += minY - height[i]
		}
	}

	return total
}


/**
方法1:
双指针法：
左右各开始一个指针，比较一下，小的向中间移动一格，同时判断当前值与前一个值的大小
如果小于，则将前一个值复制给这个值。最后算数组的总和减去原数组和即为水数。

010230102 = 011232222
9           15

15-9=6

时间复杂度：O(n)
空间复杂度：O(1)
*/
func trap2(height []int) int {
	heightLen := len(height)
	if heightLen <= 2 {
		return 0
	}

	oriSum := 0
	for _,v := range height {
		oriSum += v
	}

	l := 0
	r := heightLen - 1

	for l < r {
		if height[l] < height[r] {
			if height[l+1] < height[l] {
				height[l+1] = height[l]
			}
			l++
		} else {
			if height[r-1] < height[r] {
				height[r-1] = height[r]
			}
			r--
		}
	}

	afterSum := 0
	for _,v := range height {
		afterSum += v
	}

	return afterSum - oriSum
}
