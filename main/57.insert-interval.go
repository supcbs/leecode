package main

import "fmt"

/**

题目:
给出一个无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠
（如果有必要的话，可以合并区间）。

示例 1:
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]

示例 2:
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

*/

func main() {
	intervals := []Interval{
		{2, 5},
		{6, 7},
		{8, 9},
	}
	newInterval := Interval{0,1}
	r := insertOK(intervals, newInterval)
	fmt.Println(r)
}

/**
方法1:
遍历每个区间，需要满足以下任一条件：
a.最大小于指定区间的最小数
b.最小大于指定区间的最大数
就说明没有相交则可以直接纳入结果数组

如果相交了，则取最小的数与最后一个相交的最大的数,组成新的newInterval，
直到后续满足最小大于指定区间的最大数的时候再把这个newInterval添加到数组。

时间复杂度：O(n)
空间复杂度：O(1)
*/

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	intervalsLen := len(intervals)

	ret := make([]Interval,0)
	if intervalsLen <= 0 {
		ret = append(ret, newInterval)
		return ret
	}

	newTip := false
	for i := 0; i < intervalsLen; i++ {
		if intervals[i].Start > newInterval.End {
			if newTip == false {
				ret = append(ret, newInterval)
			}
			ret = append(ret, intervals[i])
			newTip = true
		} else if intervals[i].End < newInterval.Start {
			ret = append(ret, intervals[i])
		} else if intervals[i].End == newInterval.Start {
			newInterval.Start = intervals[i].Start
		} else if intervals[i].Start == newInterval.End {
			newInterval.End = intervals[i].End
		} else if intervals[i].End >= newInterval.Start && intervals[i].End <= newInterval.End && intervals[i].Start <= newInterval.Start{
			newInterval.Start = intervals[i].Start
		} else if intervals[i].Start >= newInterval.Start && intervals[i].Start <= newInterval.End && intervals[i].End >= newInterval.End {
			newInterval.End = intervals[i].End
		} else if intervals[i].Start <= newInterval.Start && intervals[i].End >= newInterval.End {
			newInterval.Start = intervals[i].Start
			newInterval.End = intervals[i].End
		}
	}

	if newTip == false {
		ret = append(ret, newInterval)
	}

	return ret
}

func insertOK(intervals []Interval, newInterval Interval) []Interval {
	i := 0
	resp := make([]Interval, 0)

	intervalsLen := len(intervals)

	// 比目标区间小的先加进去（end < 目标start）
	for i < intervalsLen && intervals[i].End < newInterval.Start {
		resp = append(resp, intervals[i])
		i++
	}

	// 合并目标区间（start <= 目标end）
	for i < intervalsLen && intervals[i].Start <= newInterval.End {
		newInterval.Start = Min(intervals[i].Start, newInterval.Start)
		newInterval.End = Max(intervals[i].End, newInterval.End)
		i++
	}
	resp = append(resp, newInterval)

	// 加入后续的区间
	for i < intervalsLen {
		resp = append(resp, intervals[i])
		i++
	}
	return resp
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}