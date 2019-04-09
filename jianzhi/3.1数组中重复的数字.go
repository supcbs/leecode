package main

import "fmt"

/**
题目：
有n个数组组成的数组，取值范围为0 ～ n-1，找出重复的数组

分析：
1.可以使用排序+遍历比值的方式，每次比较当前值与前一个是否相等来判断
2.可以采用map的方式，如果重复了就ok了
3.由于题目特殊，取值范围只有0～n-1，可以采用下标法进行
*/
func main() {
	arr := []int{2,3,1, 0, 2, 5, 3}
	n := 7
	r := duplicationInArray(arr, n)
	fmt.Println(r)
}

/**
1.从0开始遍历，如果当前值等于其下标，则下一个，否则就将当前值放到其下表的位置
2.如果其下标位置的数=当前值，那么结束
*/
func duplicationInArray(arr []int, n int) int {
	if len(arr) <= 0 || n != len(arr) {
		return -1
	}

	for i := 0; i < len(arr); i++ {

		for arr[i] != i {
			t := arr[i]
			if t == arr[t] {
				return t
			}

			arr[i],arr[t] = arr[t],arr[i]
		}

	}

	return -1
}