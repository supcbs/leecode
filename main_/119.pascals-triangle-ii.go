package main

import "fmt"

/**

题目:
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

输入: 3
输出: [1,3,3,1]

[
     [1],    --- 0
    [1,1],   --- 1
   [1,2,1],  --- 2
  [1,3,3,1], --- 3
 [1,4,6,4,1] --- 4
]
*/

func main() {
	rowIndex := 3
	r := getRowNew(rowIndex)
	fmt.Println(r)
}

/**
方法1:
关键点在于：
a.边界的判断，如果是0行直接返回空
b.因为只返回第n行，所以不需要像118题那样需要temp暂存直接算即可
*c.每次都将上一次的结果前头加一个1之后，当作下一次运算的开始
比如 121算完后，再加上1，为1121，
在此基础上继续算出1331，得知，第j位(非首尾位)j位加上j+1


边界注意：从0开始，所以需要<=rowIndex && 1位无法直接返回

时间复杂度：O(n^2)
空间复杂度：O(n)
*/
func getRowNew(rowIndex int) []int {
	ret := []int{1}
	if rowIndex == 0 {
		return ret
	}

	for i := 1; i <= rowIndex; i++ {
		ret = append([]int{1}, ret...)
		for j := 1; j < i; j++ {
			ret[j] = ret[j] + ret[j+1]
		}
	}

	return ret
}
