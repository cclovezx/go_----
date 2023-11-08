package Stack

import "fmt"

//定义栈类
type Stack struct {
	ElemType []int
}

//创建栈
func NewStack() *Stack {
	return &Stack{}
}

//入栈操作
func (s *Stack) Push(data int) {
	s.ElemType = append(s.ElemType, data)
}

//出栈操作
func (s *Stack) Pop() int {
	if len(s.ElemType) == 0 {
		fmt.Println("Stack is Empty!")
		return -1
	} else {
		res := s.ElemType[len(s.ElemType)-1]
		s.ElemType = s.ElemType[:len(s.ElemType)-1]
		return res
	}
}

//获取栈顶元素
func (s *Stack) Peek() int {
	if len(s.ElemType) == 0 {
		fmt.Println("Stack is Empty!")
		return -1
	} else {
		return s.ElemType[len(s.ElemType)-1]
	}
}

//查看栈是不是空
func (s *Stack) IsEmpty() bool {
	if len(s.ElemType) == 0 {
		return true
	}
	return false
}
func main() {

}
