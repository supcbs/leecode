package main

import (
	"fmt"
	"sort"
)

/**

题目:

第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
返回载到每一个人所需的最小船数。(保证每个人都能被船载)。


示例 1：

输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)

*/



func main() {
	t1 := []int{3,2,2,1} // 3
	t1 = []int{2,49,10,7,11,41,47,2,22,6,13,12,33,18,10,26,2,6,50,10} // 50
	r := numRescueBoats(t1, 50)
	fmt.Println(r)
}

/**
方法1:

没限制人员的解法


时间复杂度：O(n)
空间复杂度：O(1)
*/
func numRescueBoats(people []int, limit int) int {
	// 边界值判断
	if len(people) <= 0 || limit <= 0 {
		return 0
	}

	// 先将人员按照重量排序
	sort.Ints(people)

	// 初始化双指针的初始值
	i := 0
	j := len(people) - 1
	num := 0

	// 双指针，优先排重的
	for i <= j {
		// 题目写着只能坐两个人
		if people[j] + people[i] <= limit && i < j {
			i++
		}
		j--
		num++
	}

	return num
}

func numRescueBoats2(people []int, limit int) int {
	quickSort(people,0,len(people)-1)

	head := 0
	last := len(people)-1

	count := 0

	for last >= head {
		if people[last] + people[head] <= limit {
			count++
			last--
			head++
		} else {
			count++
			last--
		}
	}
	return count
}

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[start + (end - start)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}


/**
方法:

没限制人员的解法


时间复杂度：O(n)
空间复杂度：O(1)
*/
func numRescueBoatsWithNoLimit(people []int, limit int) int {
	// 边界值判断
	if len(people) <= 0 || limit <= 0 {
		return 0
	}

	// 先将人员按照重量排序
	sort.Ints(people)

	// 初始化双指针的初始值
	i := 0
	j := len(people) - 1
	num := 0

	// 双指针，优先排重的
	for i <= j {
		cur := people[j]
		for cur < limit {
			if cur + people[i] <= limit && i < j {
				cur += people[i]
				i++
			} else {
				break
			}
		}
		j--
		num++
	}

	return num
}
