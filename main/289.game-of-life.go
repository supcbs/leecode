package main

import "fmt"

/**

题目:
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在1970年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。
每个细胞具有一个初始状态 live（1）即为活细胞， 或 dead（0）即为死细胞。
每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上细胞的下一个（一次更新后的）状态。
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。

示例:

输入:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
进阶:

你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

*/

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}

/**
方法1:

这道题有点多文字，简化一下：

给一个矩阵，0代表死细胞，1代表活细胞
当活细胞周围的8个细胞，活细胞数等于2或者3的时候，活细胞依旧或者，否则死亡
当死细胞周围活细胞等于3个时候，死细胞复活。

本体难点在于遍历的时候，需要直到之前被修改过的细胞原先的状态，不能拿变换后的来判断。
所以假设：
die  - die  = 0
live - live = 1
live - die  = 2(这个有讲究的，这样的话，取余是一样的结果)
die  - live = 3

(注意本题只需要关注活细胞的数量)

然后再将数组遍历一把，对2取余恢复0，1


时间复杂度：O(n)
空间复杂度：O(1)
*/
func gameOfLife(board [][]int) {
	rows := len(board)
	if rows <= 0 {
		return
	}
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// 判断周围8个位置的情况
			cnt := 0
			iStart := getMaxII(0, i-1)
			iEnd := getMinII(i+1, rows-1)
			jStart := getMaxII(0, j-1)
			jEnd := getMinII(j+1, cols-1)

			for x := iStart; x <= iEnd; x++ {
				for y := jStart; y <= jEnd; y++ {
					// 自己则跳过
					if x == i && y == j {
						continue
					}

					if board[x][y] == 1 || board[x][y] == 2 {
						cnt++
					}
				}
			}

			curBoard := board[i][j]
			if (cnt < 2 || cnt > 3) && curBoard == 1 {
				board[i][j] = 2
			} else if cnt == 3 && curBoard == 0 {
				board[i][j] = 3
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Println(i,j)
			board[i][j] %= 2
		}
	}

	fmt.Println(board)

	return
}

func getMaxII(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMinII(a, b int) int {
	if a < b {
		return a
	}
	return b
}
