package main

import "fmt"

/**
关注：1.最好最差平均复杂度  2.是否稳定
*/
func main() {
	arr := []int{3, 2, 6, 4, 5, 7, 1, 8, 9}
	//r := bubbleSort(arr)
	//r := quickSort(arr, 0, len(arr)-1)
	//r := insertSort(arr)
	//r := shellSort(arr)
	//r := selectSort(arr)
	//r := heapSort(arr)
	//r := mergeSort(arr)
	r := countingSort(arr, 9)
	fmt.Println(r)
}

/**
[交换排序]冒泡排序
---
最差复杂度：O(n^2)
最好复杂度：O(n)
平均复杂度：O(n^2)

是稳定的
*/
func bubbleSort(arr []int) []int {
	if len(arr) <= 0 {
		return []int{}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

/**
1.后面已排好序了，则记录是否进行交换的标记，没交换则认为后面有序了
2.局部乱序，标记最后一个交换的位置K，下一次循环的时候，循环到K位置即可
3.依次排序确定两个值，最大和最小值，正向扫描好到最大值交换最后，反向最小放最前

from:https://blog.csdn.net/hansionz/article/details/80822494
*/
func bubbleSort1(arr []int) []int {
	if len(arr) <= 0 {
		return []int{}
	}

	k := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		pos := 0
		flag := 0
		//for j := 0; j < len(arr)-1-i; j++ {
		for j := 0; j < k; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
				pos = j
			}
		}

		if flag == 0 {
			return arr
		}

		k = pos
	}

	return arr
}

/**
[交换排序]快速排序
---
最差复杂度：O(n^2)
最好复杂度：O(n)
平均复杂度：O(nlogN)

是不稳定的
*/
func quickSort(arr []int, left int, right int) []int {
	if left < right {
		pos := partition(arr, left, right)
		quickSort(arr, left, pos-1)
		quickSort(arr, pos+1, right)
	}

	return arr
}

func partition(arr []int, left int, right int) int {
	k := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < k {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

/**
[插入排序]插入排序
---
最差复杂度：O(n^2)
最好复杂度：O(n)
平均复杂度：O(n^2)

是稳定的
*/
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			tmp := arr[i]

			j := 0
			for j = i - 1; j >= 0 && arr[j] > tmp; j-- {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
			arr[j+1] = tmp
		}
	}

	return arr
}

/**
[插入排序]希尔排序
---
最差复杂度：O(n^1.3)
最好复杂度：O(n^1.3)
平均复杂度：O(n^1.3)

是非稳定的

1.插入排序在有序的时候跑得快、2.每次只能交换一个
*/
func shellSort(arr []int) []int {
	for h := len(arr) / 2; h > 0; h = h / 2 {
		for i := h; i < len(arr); i += h {
			if arr[i] < arr[i-h] {
				tmp := arr[i]
				j := 0
				for j = i - h; j >= 0 && arr[j] > tmp && j+h < len(arr); j -= h {
					arr[j], arr[j+h] = arr[j+h], arr[j]
				}
				arr[j+h] = tmp
			}
		}

	}

	return arr
}

/**
[选择排序]选择排序
---
最差复杂度：O(n^2)
最好复杂度：O(n^2)
平均复杂度：O(n^2)

是不稳定的
*/
func selectSort(arr []int) []int{
	for i:=0;i<len(arr);i++{
		min := i
		for j:=i;j<len(arr)-1;j++{
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i],arr[min] = arr[min],arr[i]
	}

	return arr
}

/**
[选择排序]堆排序
---
最差复杂度：O(nlogN)
最好复杂度：O(nlogN)
平均复杂度：O(nlogN)

是不稳定的
*/
func heapSort(arr []int) []int{
	//造最大堆
	for i:=len(arr)/2;i>=0;i--{
		heapify(arr, i, len(arr))
	}

	//每次取出最大的，后，继续选出最大的堆
	arrLen := len(arr)
	for i:=len(arr)-1;i>0;i--{
		arr[i],arr[0]=arr[0],arr[i]
		arrLen--
		heapify(arr,0,arrLen)
	}

	return arr
}

func heapify(arr []int, i int, len int) {
	left := 2*i+1
	right := 2*i+2
	large := i

	if left<len && arr[left] > arr[large] {
		large = left
	}

	if right<len && arr[right] > arr[large] {
		large = right
	}

	if large != i {
		arr[i],arr[large] = arr[large],arr[i]
		heapify(arr,large,len)
	}

	return
}

/**
归并排序
---
最差复杂度：O(nlogN)
最好复杂度：O(nlogN)
平均复杂度：O(nlogN)

是稳定的
*/
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr)/2
	return merge(mergeSort(arr[0:mid]), mergeSort(arr[mid:]))
}

func merge(left []int, right []int) []int {
	var ret []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		ret = append(ret, left...)
	}

	if len(right) > 0 {
		ret = append(ret, right...)
	}

	return ret
}

// 基数排序


// 计数排序
/**
计数排序
---
最差复杂度：O(n+k)
最好复杂度：O(n+k)
平均复杂度：O(n+k)

是稳定的
*/
func countingSort(arr []int, maxNum int) []int {
	tmp := make([]int, maxNum+1)
	for i:=0;i<len(arr);i++{
		tmp[arr[i]]++
	}

	sortIndex :=0
	for i:=0;i<len(tmp);i++{
		for tmp[i] > 0 {
			arr[sortIndex] = i
			sortIndex++
			tmp[i]--
		}
	}

	return arr
}


// 桶排序
