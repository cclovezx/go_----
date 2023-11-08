package main

import (
	"fmt"
	Stack "workSpace/Stack/arrayStackk"
)

func main() {
	fmt.Println(tenToTwo(178923))
}

func tenToTwo(n int) int {
	stack := Stack.NewStack()
	temp := n
	if n > 1 {
		for temp > 0 {
			mod := temp % 2
			stack.Push(mod)
			temp = temp / 2
		}
		v := int(stack.Pop())
		for !stack.IsEmpty() {
			v = v*10 + stack.Pop()
		}
		return v
	} else {
		return n
	}
}
