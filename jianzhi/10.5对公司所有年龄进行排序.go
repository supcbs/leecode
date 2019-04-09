package main

import "fmt"

/**
题目：
要求效率为O(n),实现对公司的所有年龄进行排序。

分析：
先限制0-99岁的范围
先对数据进行map归类，再根据每个value的数值进行排列
*/
func main() {
	arr := []int{32,32,34,34,25,56,26,37,27,28,29}
	r := ageSort(arr)
	fmt.Println(r)
}

func ageSort(ages []int) []int {
	if len(ages) <= 0 {
		return nil
	}

	ageMaps := make(map[int]int, 99)

	for _,v:=range ages {
		if v <=99 || v > 1 {
			ageMaps[v]++
		}
	}

	ret := make([]int,0)
	for i:=1;i<=99;i++{
		if ageMaps[i] > 0 {
			for j:=1;j<=ageMaps[i];j++{
				ret = append(ret, i)
			}
		}
	}

	return ret
}