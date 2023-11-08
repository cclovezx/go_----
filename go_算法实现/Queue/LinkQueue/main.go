package Queue

import "fmt"

// 定义链表节点结构
type Node struct {
	value int
	next  *Node
}

// 定义队列结构
type Queue struct {
	head *Node
	tail *Node
}

// 初始化队列
func InitQueue() *Queue {
	return &Queue{}
}

// 入队操作
func (q *Queue) Enqueue(value int) {
	newNode := &Node{value: value, next: nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

// 出队操作
func (q *Queue) Dequeue() int {
	if q.head == nil {
		return -1 // 队列为空，返回一个错误值
	}
	item := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return item
}

func main() {
	queue := InitQueue()         // 初始化队列
	queue.Enqueue(1)             // 入队元素 1
	queue.Enqueue(2)             // 入队元素 2
	queue.Enqueue(3)             // 入队元素 3
	fmt.Println(queue.Dequeue()) // 出队元素 1，输出 1
	fmt.Println(queue.Dequeue()) // 出队元素 2，输出 2
	fmt.Println(queue.Dequeue()) // 出队元素 3，输出 3
}
