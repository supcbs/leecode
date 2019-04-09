package main

import "fmt"

/**

题目:
我们正在玩一个猜数游戏，游戏规则如下：
我从 1 到 n 之间选择一个数字，你来猜我选了哪个数字。
每次你猜错了，我都会告诉你，我选的数字比你的大了或者小了。
然而，当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。直到你猜到我选的数字，你才算赢得了这个游戏。
示例:

n = 10, 我选择了8.
第一轮: 你猜我选择的数字是5，我会告诉你，我的数字更大一些，然后你需要支付5块。
第二轮: 你猜是7，我告诉你，我的数字更大一些，你支付7块。
第三轮: 你猜是9，我告诉你，我的数字更小一些，你支付9块。

游戏结束。8 就是我选的数字。

你最终要支付 5 + 7 + 9 = 21 块钱。
给定 n ≥ 1，计算你至少需要拥有多少现金才能确保你能赢得这个游戏。

*/

func main() {
	n := 6
	r := getMoneyAmount(n)
	fmt.Println(r)
}

/**
  dp[i][j]表示从[i,j]中猜出正确数字所需要的最少花费金额.(dp[i][i] = 0)
  假设在范围[i,j]中选择x, 则选择x的最少花费金额为: max(dp[i][x-1], dp[x+1][j]) + x
  用max的原因是我们要计算最坏反馈情况下的最少花费金额(选了x之后, 正确数字落在花费更高的那侧)

  初始化为(n+2)*(n+2)数组的原因: 处理边界情况更加容易, 例如对于求解dp[1][n]时x如果等于1, 需要考虑dp[0][1](0不可能出现, dp[0][n]为0)
  而当x等于n时, 需要考虑dp[n+1][n+1](n+1也不可能出现, dp[n+1][n+1]为0)

  如何写出相应的代码更新dp矩阵, 递推式dp[i][j] = min(dp[i][j], max(dp[i][x-1], dp[x+1][j]) + x), x~[i:j], 可以画出矩阵图协助理解, 可以发现
  dp[i][x-1]始终在dp[i][j]的左部, dp[x+1][j]始终在dp[i][j]的下部, 所以更新dp矩阵时i的次序应当遵循bottom到top的规则, j则相反, 由于
  i肯定小于等于j, 所以我们只需要遍历更新矩阵的一半即可(下半矩阵)
  **/

/**
方法1:

dp[i][j] = min(dp[i][j], max(dp[i][x-1],dp[[x+1][j]) + x), x~[i:j]

时间复杂度：O(n)
空间复杂度：O(1)
*/
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)

	// 初始化为0的
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for l := 1; l < n; l++ {
		fmt.Println(n-l)
		for i := 1; i <= n-l; i++ {
			// l代表间距
			j := i + l
			dp[i][j] = 655350
			for k := i; k <= j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
		for _,b :=range dp {
			fmt.Println(b)
		}
		fmt.Println("-------")
	}

	for _,b :=range dp {
		fmt.Println(b)
	}


	return dp[1][n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
