package main

import (
	"fmt"
	"sort"
)

/**

题目:

*/

func main() {
	obj := Constructor()
	obj.AddNum(1)
	r := obj.GetIntervals()
	fmt.Println(r)
}

/**
方法1:

这个优点时间换空间了：

将范围存到两个map中，一个是大到小的映射，一个是小到大的映射。
这样做的目的是可以更快的遍历查询到是否存在val+1和val-1的范围存在。
存在的话，则进行两个映射的合并
不存在的化，创建两个映射

再将start映射的数组进行抽取int类型的排序数组存储起来，方便后续的输出数据。

时间复杂度：O(n)
空间复杂度：O(1)
*/

type Interval352 struct {
	Start int
	End   int
}

type SummaryRanges struct {
	startM map[int]int
	endM   map[int]int
	st     []int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		startM: map[int]int{},
		endM:   map[int]int{},
		st:     []int{},
	}
}

func (this *SummaryRanges) AddNum(val int) {
	startM, endM := this.startM, this.endM
	flag := false

	for start, end := range startM {
		// 说明在范围中
		if val >= start && val <= end {
			return
		} else if val == end+1 { // 范围后一个
			startM[start] = val
			delete(endM, end)
			endM[val] = start
			// 如果下一个存在的话，则进行合并
			if _, ok := startM[val+1]; ok {
				startM[start] = startM[val+1]
				endM[startM[val+1]] = start
				delete(startM, val+1)
				delete(endM, val)
			}
			flag = true
			break
		} else if val == start-1 {
			startM[val] = end
			delete(startM, start)
			endM[end] = val
			if _, ok := endM[val-1]; ok {
				startM[endM[val-1]] = end
				endM[end] = endM[val-1]
				delete(endM, val-1)
				delete(startM, val)
			}
			flag = true
			break
		}
	}

	if !flag {
		startM[val] = val
		endM[val] = val
	}

	// 变成start的排序好的数组
	s := make([]int, len(startM))

	index := 0
	for key := range startM {
		s[index] = key
		index++
	}

	sort.Ints(s)
	this.st = s
}

func (this *SummaryRanges) GetIntervals() []Interval352 {
	s := this.st
	m := this.startM
	ret := make([]Interval352, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = Interval352{s[i], m[s[i]]}
	}
	return ret
}

/**
方法2:

比较正常的解法：

1.将每个数进行排序，通过从头遍历数子比大小，得到插入的idx位置，然后插入
2.遍历得到的数字的数组，然后根据当前个一定是前一个+1，否则就直接加入范围数组
3.最后还需要处理一下第二步结尾的数据

时间复杂度：O(n)
空间复杂度：O(1)
*/
/*

type SummaryRanges struct{

	nums []int
}

func Constructor() SummaryRanges {
	var s SummaryRanges
	return s
}

func (this *SummaryRanges) AddNum(val int) {
	if len(this.nums) == 0{
		this.nums = []int{val}
		return
	}

	idx := 0

	for ;idx < len(this.nums); idx ++{
		if this.nums[idx] == val {
			return
		}

		if this.nums[idx] > val {
			break
		}
	}

	// 尾巴
	if idx == len(this.nums) {
		this.nums = append(this.nums, val)
	} else if idx == 0 { // 头部
		this.nums = append([]int{val}, this.nums...)
	} else { // 中间
		t := append([]int{val}, this.nums[idx:]...)
		this.nums = append(t, this.nums[:idx]...)
	}
}

func (this *SummaryRanges) GetIntervals() []Interval352 {
	ret := []Interval352{}

	if len(this.nums) == 0 {
		return ret
	}

	start := 0
	// 当前数不为前一个数+1，则需要如数组
	for i:=1; i< len(this.nums);i++{
		if this.nums[i] != this.nums[i-1] + 1 {
			var t Interval352
			t.Start = this.nums[start]
			t.End = this.nums[i-1]

			ret = append(ret, t)
			start = i
		}
	}

	// 最后一个扫尾
	var t Interval352
	t.Start = this.nums[start]
	t.End = this.nums[len(this.nums) - 1]
	ret = append(ret, t)
	return ret
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
*/
