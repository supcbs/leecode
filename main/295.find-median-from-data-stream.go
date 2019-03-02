package main

import (
	"fmt"
	"sort"
)

/**

题目:
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:

如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

*/

func main() {
	obj := Constructor2()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	obj.AddNum(4)
	p := obj.FindMedian()
	fmt.Println(p)
}

/**
方法1:

会超时，后续刷到最大最小堆的时候，再回头做这个

时间复杂度：O(n)
空间复杂度：O(1)
*/
type MedianFinder struct {
	nums []int
}


/** initialize your data structure here. */
func Constructor2() MedianFinder {
	return MedianFinder{
		[]int{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	this.nums = append(this.nums, num)
	sort.Ints(this.nums)
}


func (this *MedianFinder) FindMedian() float64 {
	numsLen := len(this.nums)
	var idx1,idx2 int
	if numsLen % 2 == 0 {
		idx1 = numsLen / 2
		idx2 = idx1 - 1
	} else {
		idx1 = numsLen / 2
		idx2 = idx1
	}

	return float64(this.nums[idx1] + this.nums[idx2])/2.0
}
