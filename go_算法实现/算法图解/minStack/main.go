package main

import "fmt"

//最小栈的实现

type MinStack struct {
	ElemType []int
}

// 创建栈
func NewMinStack() *MinStack {
	return &MinStack{}
}

// 判断栈是否为空
func (m *MinStack) IsEmpty() bool {
	n := len(m.ElemType)
	if n == 0 {
		return true
	}
	return false
}

// 获取栈顶元素
func (m *MinStack) Peek() int {
	if m.IsEmpty() {
		fmt.Println("Stack is Empty")
		return -1
	} else {
		return m.ElemType[len(m.ElemType)-1]
	}
}

// 压栈
func (m *MinStack) Push(data int) {
	if m.IsEmpty() {
		m.ElemType = append(m.ElemType, data)
	} else {
		if m.Peek() < data {
			max := m.Pop()
			m.ElemType = append(m.ElemType, data)
			m.ElemType = append(m.ElemType, max)
		} else {
			m.ElemType = append(m.ElemType, data)
		}
	}
}

// 弹栈
func (m *MinStack) Pop() int {
	if m.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	} else {
		top := m.Peek()
		m.ElemType = m.ElemType[0 : len(m.ElemType)-1]
		return top
	}
}

func main() {
	stack := NewMinStack()
	list := []int{2, 3, 43523, 62, 56, 457, 10, 23, 2}
	for _, i := range list {
		stack.Push(i)
	}
	fmt.Println(stack.Peek())
}
