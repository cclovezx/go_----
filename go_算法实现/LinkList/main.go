package main

import (
	"fmt"
)

//单链表创建

// 定义一个链表的结构体
type Node struct {
	value int
	Next  *Node
}

// 初始化链表
func NewNode(value int) *Node {
	return &Node{value, nil}
}

// 寻找链表尾部
func (n *Node) FindLast() *Node {
	current := n.Next
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// 头插法
func (n *Node) AddFirst(value int) {
	newNode := NewNode(value)
	newNode.Next = n.Next
	n.Next = newNode
}

// 尾插法
func (n *Node) AddLast(value int) {
	newNode := NewNode(value)
	if n.Next == nil {
		n.Next = newNode
	} else {
		n.FindLast().Next = newNode
	}
}

// 遍历链表
func (n *Node) Print() {
	current := n.Next
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.Next
	}
}

// 查找链表中的节点
func (n *Node) FindValue(value int) *Node {
	current := n.Next
	for current.Next != nil {
		if current.value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// 更新值
func (n *Node) Replace(value int, index int) {
	current := n.FindValue(index)
	current.value = value
}

// 求链表的长度
func (n *Node) Len() int {
	current := n.Next
	len := 1
	for current.Next != nil {
		len++
		current = current.Next
	}
	return len
}

// 删除节点
func (n *Node) Delete(value int) {
	current := n.Next
	for current.Next != nil {
		find := current.Next
		if find.value == value {
			break
		}
	}
	if current.Next == nil && current.value != value {
		fmt.Println("删除元素不存在")
	}
	current.Next = current.Next.Next
}
func main() {
	//创建头节点
	head := NewNode(0)
	head.AddFirst(5)
	head.AddLast(10)
	head.Print()
	len := head.Len()
	fmt.Println(len)
}
