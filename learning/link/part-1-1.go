package main

import (
	"fmt"
	_ "net/http/pprof"
)

type MyLinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: &Node{},
		Size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// 超出范围则返回-1
	if index < 0 || index >= this.Size {
		return -1
	}

	cur := this.Head
	// 这里要注意 i=0 的时候需要返回第一个
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}

	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Val:  val,
		Next: this.Head.Next, // 这里有点坑，题目定义的是第一个值为0，next才是正式开始的第一个值
	}

	this.Head.Next = newNode
	this.Size++
}

/**
index   0    1    2
head -> 1 -> 2 -> 3 -> nil
*/

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Size == 0 {
		this.AddAtHead(val)
		return
	}

	pred := this.Head
	// 只要下一个一直存在就ok
	for pred.Next != nil {
		pred = pred.Next
	}

	pred.Next = &Node{
		Val:  val,
		Next: nil,
	}
	this.Size++
}

/**
index   0    1    2
head -> 1 -> 2 -> 3 -> nil
*/

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 边界值处理。如果不再1-1000则返回，如果不在已存在的长度则返回
	if index > this.Size {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	pred := this.Head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}

	newNode := &Node{
		Val:  val,
		Next: pred.Next,
	}

	pred.Next = newNode
	this.Size++
}

/**
head -> 1 -> 2 -> 3 -> nil
*/
/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	}

	pred := this.Head
	// 这里index需要减去1，这样next代表要删除的元素
	for i := 0; i < index; i++ {
		pred = pred.Next
	}

	pred.Next = pred.Next.Next
	this.Size--
}

func (this *MyLinkedList)Print(msg string){
	fmt.Printf("%s: [", msg)
	cur := this.Head
	for cur!=nil {
		fmt.Printf("%v\t", cur.Val)
		cur = cur.Next
	}
	fmt.Println("]")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	linklist := Constructor()
	linklist.AddAtHead(1)
	linklist.AddAtTail(3)
	linklist.AddAtIndex(1,2)   //链表变为1-> 2-> 3
	v := linklist.Get(1)            //返回2
	fmt.Println(v)
	linklist.DeleteAtIndex(0)  //现在链表是1-> 3
	v = linklist.Get(0)
	fmt.Println(v)
}


/**
经验：
1.这里注意index是从0开始的
2.获取的index如果=size也是获取不到的
3.如果用cur当作标杆则for循环index需要+1，如果是pred则for存还不用+1
 */
