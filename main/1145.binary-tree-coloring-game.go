package main

import "fmt"

/**
题目:
有两位极客玩家参与了一场「二叉树着色」的游戏。游戏中，给出二叉树的根节点 root，树上总共有 n 个节点，且 n 为奇数，其中每个节点上的值从 1 到 n 各不相同。
游戏从「一号」玩家开始（「一号」玩家为红色，「二号」玩家为蓝色），最开始时，
「一号」玩家从 [1, n] 中取一个值 x（1 <= x <= n）；
「二号」玩家也从 [1, n] 中取一个值 y（1 <= y <= n）且 y != x。
「一号」玩家给值为 x 的节点染上红色，而「二号」玩家给值为 y 的节点染上蓝色。

之后两位玩家轮流进行操作，每一回合，玩家选择一个他之前涂好颜色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色。
如果当前玩家无法找到这样的节点来染色时，他的回合就会被跳过。
若两个玩家都没有可以染色的节点时，游戏结束。着色节点最多的那位玩家获得胜利 ✌️。
现在，假设你是「二号」玩家，根据所给出的输入，假如存在一个 y 值可以确保你赢得这场游戏，则返回 true；若无法获胜，就请返回 false。

示例：
输入：root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
输出：True
解释：第二个玩家可以选择值为 2 的节点。
 
提示:
二叉树的根节点为 root，树上由 n 个节点，节点上的值从 1 到 n 各不相同。
n 为奇数。
1 <= x <= n <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = nums[0]

	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]

	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}

			ch <- tree.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}

			ch <- tree.Right
		}
	}

	return root
}

func main() {
	t1 := buildTree([]int{2,1,4})
	r := btreeGameWinningMove(t1,1,2)
	fmt.Println(r)
}

var l,r int
func dfs(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, x)
	right := dfs(root.Right, x)
	if root.Val == x {
		l = left
		r = right
	}

	return 1 + left + right
}

/**
方法：
这道题核心在于分析问题：
1.着色只能是已选节点旁边节点（父节点，左节点，右节点）
2.先手选择一个节点后，剩下的部分分为三个部分（父节点关联部分，左节点关联部分，右节点关联部分），
  后手只能选择三个部分其一才会是最优解，因为这样占据的面积才更大
3.问题就转化为，求解这三个部分存不存在数量大于n/2的，存在的话，就后手赢，否则先手

考查知识点：深度遍历（dfs）
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	// 初始化节点
	dfs(root, x)
	p := n - l - r - 1
	if l > n/2 || r > n/2 || p > n/2 {
		return true
	}

	return false
}
