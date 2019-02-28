package main

import "fmt"

/**

题目:
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
你需要按照以下要求，帮助老师给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:
输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。

示例 2:
输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。

*/

func main() {
	nums := []int{1,0,2}
	r := candy(nums)
	fmt.Println(r)
}

/**
方法1:

思路很简单，首先每个人至少拿到1个，
感觉和42题有一丁点像。都是从左往右，再从右往左

从左往右遍历一遍，如果num[i]>num[i-1],那么num[i] = num[i-1]+1
从右往左遍历一遍，如果num[i]>num[i+1] && map[i]<=map[i+1],那么num[i] = num[i+!]+1

需要假设越来越大，所以只能和遍历方向的后头比

时间复杂度：O(3n)
空间复杂度：O(n)
*/
func candy(ratings []int) int {
	ratingsLen := len(ratings)
	if ratingsLen <= 0 {
		return 0
	}

	candies := make(map[int]int, ratingsLen)
	for i := 0; i < ratingsLen; i++ {
		candies[i] = 1
	}

	for i := 1; i < ratingsLen; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := ratingsLen - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	sum := 0
	for _, v := range candies {
		sum += v
	}

	return sum
}
