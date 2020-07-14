package main

import "fmt"

type MyLinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
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
		Prev: this.Head,
	}

	if this.Head.Next != nil {
		this.Head.Next.Prev = newNode
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
		Prev: pred,
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
		Prev: pred,
	}

	// 如果存在下一个节点，那么需要重置一下pred
	if pred.Next != nil {
		pred.Next.Prev = newNode
	}

	pred.Next = newNode
	this.Size++
}

/**
index   0    1    2
head -> 1 -> 2 -> 3 -> nil
*/

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	}

	pred := this.Head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}

	// 如果其下一个的下一个存在值，那么需要重置一些pred
	if pred.Next.Next != nil {
		pred.Next.Next.Prev = pred.Next
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

