package main

import "fmt"

/**
TOP N的问题编程。

from：https://www.jianshu.com/p/eb28128bd266

分析：核心思路在于内存和机器条件
1.内存足够，可以直接使用排序，采用快排的方式，以一个数作为基准，进行划分，大的部分如果大于N个数量，则进行继续在大的中寻找，小的一样，直到正好找到N个
2.内存不够的话，如果机器够多，可以采用mapreduce的当时，将数据划分成内存中足够放的范围，每台机器进行计算top1000，然后汇总进行归并排序算出最后的topN
3.机器和内存都有限的话，那就采用最小堆的方式进行，复杂度O(nlogn)
*/

func main() {
	data := []int{2, 1, 4, 5, 2, 6, 2, 5, 7, 8, 3}

	data = buildHeap(5, data)
	for i := 5; i < len(data); i++ {
		adjust(i, 5, data)
	}

	fmt.Println(data)
}

func heapParent(n int) int {
	return (n - 1) / 2
}

func heapLeft(n int) int {
	return 2*n + 1
}

func heapRight(n int) int {
	return 2*n + 2
}

func buildHeap(n int, data []int) []int {
	if n <= 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		t := i

		// 除了第一个外，每个都与其父亲进行比较
		for t != 0 && data[heapParent(t)] > data[t] {
			data[heapParent(t)], data[t] = data[t], data[heapParent(t)]
			t = heapParent(t)
		}
	}

	return data
}

func adjust(i, n int, data []int) []int {
	if data[i] <= data[0] {
		return data
	}

	data[i], data[0] = data[0], data[i]

	t:=0
	for (heapLeft(t) < n && data[t] > data[heapLeft(t)]) ||
		(heapRight(t) < n && data[t] > data[heapRight(t)]) {
		if heapLeft(t) < n && data[t] > data[heapLeft(t)] {
			data[t],data[heapLeft(t)] = data[heapLeft(t)],data[t]
			t = heapLeft(t)
		} else {
			data[t],data[heapRight(t)] = data[heapRight(t)],data[t]
			t = heapRight(t)
		}
	}

	return data
}
