package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type CircularLinkedList struct {
	head *Node
}

//初始化节点
func NewNode(data int) *Node {
	return &Node{data, nil}
}

//头插法
func (cll *CircularLinkedList) AddFirst(data int) {
	newNode := NewNode(data)
	if cll.head == nil {
		cll.head = newNode
		newNode.next = cll.head
	} else {
		newNode.next = cll.head
		cll.head = newNode
	}
}

//尾插法
func (cll *CircularLinkedList) AddLast(data int) {
	newNode := NewNode(data)

	if cll.head == nil {
		cll.head = newNode
		newNode.next = cll.head
	} else {
		current := cll.head
		for current.next != cll.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = cll.head
	}
}

//查找
func (cll *CircularLinkedList) FindValue(data int) *Node {
	current := cll.head
	for current != cll.head && current.data != data {
		current = current.next
	}
	if current == cll.head {
		return nil
	}
	return current
}

//删除
func (cll *CircularLinkedList) Delete(data int) {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if cll.head.data == data { // 删除的是头节点
		currentNode := cll.head
		for currentNode.next != cll.head {
			currentNode = currentNode.next
		}
		currentNode.next = cll.head.next
		cll.head = cll.head.next
	} else {
		// 删除的是中间节点或尾节点
		currentNode := cll.head
		for currentNode != cll.head && currentNode.data != data {
			currentNode = currentNode.next
		}
		if currentNode == cll.head {
			fmt.Println("Item not found in the list")
			return
		} else if currentNode.next == cll.head {
			// 删除的是尾节点
			currentNode = currentNode.next
			cll.head = currentNode
		} else {
			// 删除的是中间节点
			nextNode := currentNode.next
			prevNode := currentNode.next.next
			prevNode.next = nextNode.next
			currentNode.next = cll.head
			cll.head = currentNode
		}
	}
}

//修改值
func (cll *CircularLinkedList) Replace(value int, index int) {
	current := cll.FindValue(index)
	current.data = value
}

//遍历
func (cll *CircularLinkedList) Display() {
	if cll.head == nil {
		fmt.Println("List is empty")
	} else {
		currentNode := cll.head
		for true {
			fmt.Println(currentNode.data)
			currentNode = currentNode.next
			if currentNode == cll.head {
				break
			}
		}
	}
}
func main() {
	var cll CircularLinkedList
	//cll.AddFirst(1)
	cll.AddFirst(2)
	cll.AddLast(4)
	cll.Display()
	fmt.Println(cll.FindValue(2))
}
