package main

import (
	"fmt"
)

/**

题目:
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明:
如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。

示例 1:
输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

示例 2:
输入:
gas  = [2,3,4]
cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。

*/

func main() {
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	r := canCompleteCircuitNew(gas, cost)
	fmt.Println(r)
}

func canCompleteCircuitNew(gas []int, cost []int) int {
	gasLen := len(gas)
	costLen := len(cost)

	if gasLen <= 0 || costLen <= 0 {
		return -1
	}

	var start,total,rest int
	for i := 0;i < gasLen; i++ {
		diff := gas[i] - cost[i]
		total += diff
		rest += diff
		if rest < 0 {
			start = i+1
			rest = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}



/**
跑得最快的解
 */
func canCompleteCircuitFast(gas []int, cost []int) int {
	remains, debts, start := 0, 0, 0

	for i, g := range gas {
		remains += g - cost[i]
		if remains < 0 {
			start = i + 1
			debts += remains
			remains = 0
		}
	}

	if debts+remains < 0 {
		return -1
	}
	return start
}

/**
方法：

这道题和55题最大子串是一个道理
 */

/**
方法1:
能跑完全程有两个关键点：
1.需要总的油量 > 每站损耗的
2.如果当前k站获取的油量不足以支撑到下一站，则目标站定为k+1站，继续进行判断

时间复杂度：O(n)
空间复杂度：O(1)

---
想了想能不能不用totol发现是不行的，
[7,1,0,11,4]
[5,9,1,2,5]
这种场景会报错，所以还是需要从目标站开始累计算一下油
*/
func canCompleteCircuit(gas []int, cost []int) int {
	var rest, start, total int

	gasLen := len(gas)
	for i := 0; i < gasLen; i++ {
		diffGas := gas[i] - cost[i]

		// 当前位置累计油量
		total += diffGas

		// 全程油量
		rest += diffGas

		if total < 0 {
			start = i + 1
			total = 0
		}
	}

	if rest < 0{
		return -1
	}

	return start % gasLen
}
