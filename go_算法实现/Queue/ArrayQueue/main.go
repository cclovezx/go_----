package Queue

import (
	"fmt"
)

type Queue struct {
	ElemType []int
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	if len(q.ElemType) == 0 {
		return true
	}
	return false
}

// 入队
func (q *Queue) Enqueue(data int) {
	q.ElemType = append(q.ElemType, data)
}

// 出队
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return -1
	} else {
		front := q.ElemType[0]
		q.ElemType = q.ElemType[1:]
		return front
	}
}

func main() {
	queue := Queue{}
	arr := []int{1, 3, 42, 534, 56, 4}
	for _, i := range arr {
		queue.Enqueue(i)
	}
	fmt.Println(queue.Dequeue())
}
