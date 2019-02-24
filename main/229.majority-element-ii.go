package main

import "fmt"

/**

题目:
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

示例 1:
输入: [3,2,3]
输出: [3]

示例 2:
输入: [1,1,1,3,3,2,2,2]
输出: [1,2]
*/

func main() {
	nums := []int{1,2}
	r := majorityElementii(nums)
	fmt.Println(r)
}

/**
方法1:
由于是至少大于n/3,所以至多只能有0、1、2三种可能个数。
可以使用投票法进行筛选，同169的思路，只不过此处的需要选择top2。
先假定两个数为top2，遇到相同的就+1，不同的就-1，当count为0的时候，则指定当前遍历到的数为目标数
然后选出top2后在进行计数验证一把是不是真的满足n/3

时间复杂度：O(n)
空间复杂度：O(1)
*/
func majorityElementii(nums []int) []int {
	var targetNum1, targetNum2, count1, count2 int

	// 投票选出top2
	for _,v :=range nums {
		fmt.Println(v)
		if v == targetNum1 {
			count1 ++
		} else if v == targetNum2 {
			count2 ++
		} else if count1 == 0 {
			targetNum1 = v
			count1 = 1
		} else if count2 == 0 {
			targetNum2 = v
			count2 = 1
		} else {
			count1 --
			count2 --
		}
	}

	// 判断top2是不是真的>n/3
	count1,count2 = 0,0
	for _,v :=range nums {
		if v == targetNum1 {
			count1 ++
		} else if v == targetNum2 {
			count2 ++
		}
	}

	ret := make([]int,0)
	if count1 > len(nums)/3 {
		ret = append(ret, targetNum1)
	}

	if count2 > len(nums)/3 {
		ret = append(ret, targetNum2)
	}

	return ret
}


/**
扩展：
如果题目没限定空间复杂度为 O(1)，那么可以存一个key为当前数，value为出现次数的map
然后遍历这个map得出大于n/3的数
 */