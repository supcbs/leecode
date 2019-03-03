package main

import "fmt"

/**

题目:
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

*/

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	m := 3
	n := 3
	merge88FAST(nums1, m, nums2, n)
	fmt.Println(nums1)
}

/**
方法1:

从尾部向前移动 + 双指针。
因为nums1肯定比nums2长，所以计算出合并后总的长度。
然后分别总尾巴遍历nums1和nums2，由于排好序了，所以大的肯定是最后一个。
直到其中一个遍历完成。此时会有两种情况。
1.nums1没遍历好，那么则忽略，因为它本来就在nums1里头了
2.nums2没遍历好，那说明上下的是最小的了，直接按n的排位直接替换到nums2即可

---
这里必须要说下为什么不能从前往后遍历，因为这样的化，当前头有元素插入后，后头的元素将全部移位。


时间复杂度：O(m+n)
空间复杂度：O(1)
*/
func merge88(nums1 []int, m int, nums2 []int, n int) {
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	// 空值判断
	if nums1Len <= 0 || nums2Len <= 0 {
		return
	}

	tail := m + n - 1
	m--
	n--
	// 肯定是n先遍历完，或者一起遍历完
	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[tail] = nums1[m]
			m--
		} else {
			nums1[tail] = nums2[n]
			n--
		}
		tail--
	}

	// 如果n都比m的小，那么m遍历完了，n还有，则需要放到nums1开头
	for n >= 0 {
		nums1[tail] = nums2[n]
		n--
		tail--
	}

	return
}

/**
方法2:

看了跑得最快的方式：
先将nums2拷贝到nums10的位置，然后在进行冒泡排序

时间复杂度：O(m+n)
空间复杂度：O(1)
*/
func merge88FAST(nums1 []int, m int, nums2 []int, n int) {
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	// 空值判断
	if nums1Len <= 0 || nums2Len <= 0 {
		return
	}

	copy(nums1[m:], nums2)

	fmt.Println(nums1)
	totalLen := m + n
	for i := 0; i < totalLen; i++ {
		for j := 0; j < totalLen-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}

	return
}
