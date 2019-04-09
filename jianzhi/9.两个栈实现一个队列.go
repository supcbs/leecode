package main

/**
题目：
两个栈实现一个队列

分析：
add的时候，直接往1中塞。
pop的时候，先判断2是否存在元素，存在的话直接吐出，不存在的话，将1往2中导入
*/
func main() {

}

type QueueWithTwoStack struct {
	stack1 []int
	stack2 []int
}

func (q QueueWithTwoStack) add (num int) {
	q.stack1 = append(q.stack1, num)
}

func (q QueueWithTwoStack) pop () int {
	if len(q.stack2) <= 0 {
		for len(q.stack1) > 0 {
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
		}
	}

	if len(q.stack2) < 0 {
		return -1
	}

	return q.stack2[len(q.stack2)]
}