package main

import "fmt"

/**
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

案例 1:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

输出: True


案例 2:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

输出: False
*/
func main() {
	t := []int{1, -1, 2, 2}
	d := buildTree(t)
	k := 3
	r := findTarget(d, k)
	fmt.Println(r)
}

var arr []int

func findTarget(node *TreeNode, k int) bool {
	arr = []int{}

	// 根据中序遍历获取到递增的数据
	med(node)

	// 双指针快速定位
	left := 0
	right := len(arr) - 1
	for left < right {
		// 如果两数相加大于k，则缩小最右边的值
		if arr[left]+arr[right] > k {
			right--
		} else if arr[left]+arr[right] < k {
			left++
		} else {
			return true
		}
	}

	return false
}

func med(node *TreeNode) {
	if node == nil {
		return
	}
	med(node.Left)
	arr = append(arr, node.Val)
	med(node.Right)
	return
}
