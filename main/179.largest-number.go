package main

import (
	"fmt"
	"strconv"
	"sort"
)

/**

题目:
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:
输入: [10,2]
输出: 210

示例 2:
输入: [3,30,34,5,9]
输出: 9534330

说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

*/

func main() {
	nums := []int{3,30,34,5,9}
	r := largestNumberOK(nums)
	fmt.Println(r)
}

/**
方法1:

相当于来一个冒泡排序。
只不过比对的规则是a+b>b+a即可升级

时间复杂度：O(n)
空间复杂度：O(1)
*/
func largestNumber(nums []int) string {
	numsLen := len(nums)
	if numsLen <= 0 {
		return ""
	}

	fmt.Println("0" == "00")

	strs := make([]string, len(nums))
	for i:=0;i<numsLen;i++{
		strs[i] = strconv.Itoa(nums[i])
	}
	fmt.Println(strs)
	for i:=0;i<numsLen;i++{
		for j:=0;j<numsLen-i-1;j++{
			fmt.Println(strs[j],strs[j+1])
			if strs[j] + strs[j+1] < strs[j+1] + strs[j] {
				strs[j],strs[j+1] = strs[j+1],strs[j]
			}
			fmt.Println(strs)
		}
	}

	fmt.Println(strs)

	// 拼接字符串
	ret := ""
	for _,v := range strs{
		ret += v
	}

	fmt.Println(ret[0])
	if ret[0] == '0' && numsLen >0 {
		return "0"
	}
	return ret
}


func largestNumberOK(nums []int) string {
	ss := make([]string, len(nums))
	for i := range nums {
		ss[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i] + ss[j] > ss[j] + ss[i] {
			return true
		}
		return false
	})
sort.Ints()
	fmt.Println(ss)
	var res string
	for i := range ss {
		res += ss[i]
	}
	if len(res) > 0 && res[0] == '0' {
		return "0"
	}
	return res
}
