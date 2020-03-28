package main

import "fmt"

/**
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
func main() {
	t := []int{1, -1, 2}
	d := buildTree(t)
	r := findMode(d)
	fmt.Println(r)
}

func rangeSumBST(node *TreeNode, L int, R int) (ret int) {
	if node == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	cur := node
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if cur.Val >= L && cur.Val <= R {
				ret += cur.Val
			} else if cur.Val > R {
				return
			}

			cur = cur.Right
			println(cur)
		}
	}
	return
}


var sum int
func rangeSumBSTOK(node *TreeNode, L int, R int) (ret int) {
	sum = 0
	BTSValue(node,L,R)
	return sum
}

func BTSValue(root *TreeNode,L,R int){
	if root == nil {
		return
	}

	if root.Val > L && root.Left != nil {
		BTSValue(root.Left,L,R)
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Right != nil && root.Val < R{
		BTSValue(root.Right,L,R)
	}
}
