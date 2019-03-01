package main

import (
	"fmt"
	"sort"
)

/**

题目:
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

*/

type Interval56 struct {
	Start int
	End   int
}

func main() {
	intervals := []Interval56{


		{1, 4},
		{1,2 },
		{0,0 },
	}
	r := merge(intervals)
	fmt.Println(r)
}

/**
方法1:

暂存一个temp的范围，一开始为第一个范围，
如果下一个区间和当前temp没有交集，则temp入返回队列，将当前范围设置为temp继续后面操作

如果有交集则置为更新当前temp

时间复杂度：O(n)
空间复杂度：O(1)
*/

type IL []Interval56
func (list IL) Len() int {
	return len(list)
}

func (list IL) Swap(i,j int) {
	list[i],list[j] = list[j],list[i]
}

func (list IL) Less(i,j int) bool {
	if list[i].Start == list[j].Start {
		return list[i].End < list[j].End
	}

	return list[i].Start < list[j].Start
}

func merge(intervals []Interval56) []Interval56 {
	intervalsLen := len(intervals)

	ret := make([]Interval56,0)
	if intervalsLen <= 0 {
		return ret
	}

	var il IL
	il = intervals
	sort.Sort(il)

	fmt.Println(il,intervals)

	var curInterval Interval56
	for i:=0;i<intervalsLen;i++{
		if i == 0 {
			curInterval = il[0]
			continue
		}

		// 说明没有交集
		if curInterval.End < il[i].Start {
			ret = append(ret, curInterval)
			curInterval = il[i]
			continue
		}

		if curInterval.Start > il[i].Start {
			curInterval.Start = il[i].Start
		}

		if curInterval.End < il[i].End {
			curInterval.End = il[i].End
		}
	}

	ret = append(ret, curInterval)

	return ret
}
