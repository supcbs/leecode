package main

import (
	"fmt"
	"strconv"
)

/**

我们从二叉树的根节点 root 开始进行深度优先搜索。
在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），
然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
如果节点只有一个子节点，那么保证该子节点为左子节点。
给出遍历输出 S，还原树并返回其根节点 root。



示例 1：

输入："1-2--3--4-5--6--7"
输出：[1,2,5,3,4,6,7]
示例 2：

输入："1-2--3---4-5--6---7"
输出：[1,2,5,3,null,6,null,4,null,7]
示例 3：

输入："1-401--349---90--88"
输出：[1,401,null,349,88,90]

提示：

原始树中的节点数介于 1 和 1000 之间。
每个节点的值介于 1 和 10 ^ 9 之间。
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
	str := "1-2--3--4-5--6--7"
	r := recoverFromPreorder(str)
	fmt.Println(r)
}

/**
方法1:

前序遍历进行判定

时间复杂度：O(n)
空间复杂度：O(h)
*/

func recoverFromPreorder(S string) *TreeNode {
	// 用于记录深度，预先加一个0代表root元素
	deps := []int{0}
	dep := 0

	// 可能存在>9的数字
	// 所以需要累计数字直到遇到-才能认为这个数字结束
	numByte := make([]byte, 0)

	// 用于记录遇到的数字
	nums := make([]int, 0)

	// 从字符串的左边开始遍历
	for i := 0; i < len(S); i++ {
		// fmt.Println(i, S[i] == '-', dep, string(numByte), nums)
		// "1-2--3--4-5--6--7"
		// 如果遇到'-'，说明前一个数字已经遍历结束
		if S[i] == '-' {
			// 如果层次变为0，则说明前文刚刚结束一个数字的收集
			// 在这里需要将数字添加到nums中
			if dep == 0 {
				num, _ := strconv.Atoi(string(numByte))
				nums = append(nums, num)
				numByte = make([]byte, 0)
			}

			// 继续累计深度
			dep++
		} else {
			// 如果深度部位0，说明前文刚刚算出接下来的数字的层深度
			// 此时先讲层深度记录到deps中，并且初始化dep
			if dep != 0 {
				deps = append(deps, dep)
				dep = 0
			}

			// 累计连续遇到的数字
			numByte = append(numByte, S[i])
		}
	}

	// 处理最后一个数字
	num, _ := strconv.Atoi(string(numByte))
	nums = append(nums, num)

	// 此时得到了两个数组
	// 1. 每个数字的数组
	// 2. 每个数字对应的深度的数组
	return recoverTree(deps, nums, 0, len(deps)-1)
}

/**
这里需要非常熟悉前序遍历排列出来的数字的规律：
字符串：1-2--3--4-5--6--7
数组1：[1,2,3,4,5,6,7]
数组2：[0,1,2,2,1,2,2]
         |-----|
两个相同层次之间的数字一定是属于第一个低层次的那个节点的自=子节点
*/
func recoverTree(deps []int, nums []int, start int, end int) *TreeNode {
	head := &TreeNode{
		Val:   nums[start],
		Left:  nil,
		Right: nil,
	}

	dep := deps[start]
	left := -1
	right := -1

	// 计算得到该层次的两个节点的位置（一定只有两个）
	for i := start + 1; i <= end; i++ {
		if deps[i] == dep+1 {
			if left == -1 {
				left = i
			} else {
				right = i
			}
		}
	}

	// 如果左节点都没有，那说明是nil了
	if left != -1 {
		// 如果存在右节点，则继续进行范围递归
		if right != -1 {
			// 注意右边是right-1，因为不能包括右边的节点
			head.Left = recoverTree(deps, nums, left, right-1)
			head.Right = recoverTree(deps, nums, right, end)
		} else {
			// 能到这里说明到尾巴都是左节点的了
			head.Left = recoverTree(deps, nums, left, end)
		}
	}
	return head
}

func recoverFromPreorder(S string) *TreeNode {

	l := len(S)
	if l == 0 {
		return nil
	}

	dummy := &TreeNode{}
	deepMap := make([]*TreeNode, 0)

	for pre, cur := 0, 0; cur < l; cur++ {

		if S[cur] != '-' {

			deep := cur - pre
			pre = cur

			for ; cur <= l; cur++ {

				//cur == l 防止最后一个数漏掉
				if cur == l || S[cur] == '-' {

					node := &TreeNode{}
					node.Val, _ = strconv.Atoi(S[pre:cur])

					if deep == 0 {
						dummy.Left = node
						deepMap = append(deepMap, node)
					} else {
						parent := deepMap[deep-1]
						if parent.Left == nil {
							parent.Left = node
						} else {
							parent.Right = node
						}

						//更新栈
						deepMap = deepMap[:deep]
						deepMap = append(deepMap, node)
					}

					pre = cur
					break
				}
			}
		}
	}

	return dummy.Left
}
