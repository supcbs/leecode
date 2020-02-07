package main

import "fmt"

/**

题目:
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func main() {
	numRows := 5
	r := generateNew(numRows)
	fmt.Println(r)
}

func generateNew(numRows int) [][]int {
	ret := [][]int{}
	if numRows <= 0 {
		return ret
	}

	ret = [][]int{{1}}
	if numRows == 1 {
		return ret
	}

	for i:=1;i<numRows;i++{
		tmp := make([]int, i+1)
		tmp[0] = 1
		for j:=1;j<i;j++ {
			tmp[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		tmp[i] = 1
		ret = append(ret, tmp)
	}
	return ret
}




/**
方法1:
关键点在于：
a.边界的判断，如果是0行直接返回空，1行也可以直接返回
b.首位肯定都是1，其余的数为i-1行的j-1和j位数相加而成

时间复杂度：O(n^2)
空间复杂度：O(n^2)
*/
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	ret := [][]int{{1}}
	if numRows == 1 {
		return ret
	}

	for i := 1; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		for j := 1; j < i; j++ {
			temp[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		temp[i] = 1
		ret = append(ret, temp)
	}

	return ret
}
