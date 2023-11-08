package DLinkList

import "fmt"

//节点类
type DNode struct {
	data        interface{}
	prior, next *DNode
}

//初始化头节点
type DoubleLinkList struct {
	head *DNode
	tail *DNode
}

//初始化节点
func NewNode(data int) *DNode {
	return &DNode{data, nil, nil}
}

//头插法
func (d *DoubleLinkList) AddFirst(data int) {
	node := NewNode(data)
	if d.head != nil {
		node.prior = d.head
		node.next = d.head.next
		d.head.next = node
	} else {
		d.head = node
		d.tail = node
	}
}

//尾插法
func (d *DoubleLinkList) AddLast(data int) {
	node := NewNode(data)
	current := d.head
	if current == nil {
		d.head = node
		d.tail = node
	} else {
		node.prior = d.tail
		d.tail.next = node
		d.tail = node
	}
}

//查找结点
func (d *DoubleLinkList) FindValue(data int) bool {
	current := d.head
	if current.next == nil {
		return false
	}
	for current.next != nil {
		if current.data == data {
			return true
		}
	}
	return false
}

//遍历链表
func (d *DoubleLinkList) Display() {
	current := d.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
}

func main() {
	d := DoubleLinkList{}
	d.AddFirst(1)
	d.AddLast(2)
	d.Display()
}
