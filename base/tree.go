package main

import (
	"fmt"
)

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

	for i:=0;i<len(nums);i++{
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val:nums[i],
			}

			ch <- tree.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val:nums[i],
			}

			ch <- tree.Right
		}
	}

	return root
}

func main() {
	t := buildTree([]int{1,2,3,4,5,6,7})
	//r := preOrderD(t)
	//r:=preOrder(t)
	//r := inOrderD(t)
	//r:=inOrder(t)
	//r := postOrderD(t)
	//r:=postOrder(t)

	r:=breadthFirst(t)
	fmt.Println(r)
}

// 二叉树前序遍历 - 递归
var ret []int
func preOrderD(node *TreeNode) []int {
	preOrderHelp(node)
	return ret
}
func preOrderHelp(node *TreeNode) {
	if node == nil {
		return
	}

	ret = append(ret, node.Val)
	preOrderHelp(node.Left)
	preOrderHelp(node.Right)
}

// 二叉树前序遍历 - 非递归
func preOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	var stack []*TreeNode
	stack = append(stack, node)

	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		ret = append(ret, tmp.Val)

		// 这里是right优先，因为是栈
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}

		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
	}

	return ret
}


// 二叉树中序遍历 - 递归
func inOrderD(node *TreeNode) []int {
	inOrderHelp(node)
	return ret
}

func inOrderHelp(node *TreeNode) {
	if node == nil {
		return
	}

	inOrderHelp(node.Left)
	ret = append(ret, node.Val)
	inOrderHelp(node.Right)
}

// 二叉树中序遍历 - 非递归
func inOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	stack := make([]*TreeNode,0)
	cur := node
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack,cur)
			cur = cur.Left
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
		}
	}

	return ret
}

// 二叉树后序遍历 - 递归
func postOrderD(node *TreeNode) []int {
	postOrderHelp(node)
	return ret
}

func postOrderHelp(node *TreeNode) {
	if node == nil {
		return
	}

	postOrderHelp(node.Left)
	postOrderHelp(node.Right)
	ret = append(ret, node.Val)
}

// 二叉树后序遍历 - 非递归
func postOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	var pre *TreeNode
	stack := make([]*TreeNode, 0)
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Right == nil || cur.Right == pre {
			ret = append(ret, cur.Val)
			pre = cur
		} else {
			stack = append(stack,cur)
			cur = cur.Right
			for cur != nil {
				stack = append(stack,cur)
				cur = cur.Left
			}
		}
	}

	return ret
}

// 二叉树广度优先遍历
func breadthFirst(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	queue := make([]*TreeNode,0)
	queue = append(queue, node)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		ret = append(ret,cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return ret
}


// 二叉树深度优先遍历-前序遍历


// 深度 dfs
func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(getDepth(node.Left),getDepth(node.Right)) + 1
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//最小公共父节点
/**

TREE* CommonFather(TREE *root, TREE *A, TREE *B)
{
    if(root == NULL)
        return root;
    if(root == A)//如果找到A，则后面的都不再找了，如果其他分支没找到B，则B必定在A下面
        return A;
    if(root == B)//同上
        return B;
    TREE *leftChild == NULL;
    TREE *rightChild == NULL;
    leftChild = CommonFather(root->left, A, B);//返回A，B或结果
    rightChild = CommonFather(root->right, A, B);//返回A，B或结果
    if(leftChild != NULL && rightChild != NULL)//如果都不为空，则必定一个是A，一个是B；
        return root;
    if(leftChild != NULL)//如果不为空，则必定是A或B或结果；
        return leftChild;
    if(rightChild != NULL)
        return rightChild;//如果不为空，则必定是A或B或结果；
}
 */


//两个节点距离
//两节点最长的距离

//https://blog.csdn.net/Charliewolf/article/details/81564211

