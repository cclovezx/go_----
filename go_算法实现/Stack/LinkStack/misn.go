package Stack

import "fmt"

//节点类
type StackNode struct {
	data int
	next *StackNode
}

//栈类型
type Stack struct {
	top  *StackNode
	size int
}

//初始化节点
func NewStackNode(data int) *StackNode {
	return &StackNode{data, nil}
}

//初始化栈
func (s *Stack) init() {
	s.top = nil
	s.size = 0
}

//判断链栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

//入栈
func (s *Stack) Push(data int) {
	node := NewStackNode(data)
	if s.IsEmpty() {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}
	s.size++
}

//出栈
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("Stack is None")
		return -1, false
	} else {
		topNode := s.top
		s.top = s.top.next
		topData := topNode.data
		s.size--
		return topData, true
	}
}

//获取栈顶元素
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty")
		return -1
	} else {
		return s.top.data
	}
}

func main() {

}
