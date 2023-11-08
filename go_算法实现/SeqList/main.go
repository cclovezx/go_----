package SeqList

import "fmt"

//线性表
type SeqList struct {
	ElemType []int
}

// 给线性表添加元素
func (list *SeqList) Add(element int) {
	list.ElemType = append(list.ElemType, element)
}

// 删除线性表元素(根据下标)
func (list *SeqList) RemoveList(index int) {
	if index > 0 || index > len(list.ElemType) {
		return
	}
	list.ElemType = append(list.ElemType[:index], list.ElemType[index+1:]...)
}

// 查找线性表元素
func (list *SeqList) GetIndex(index int) int {
	if index < 0 || index > len(list.ElemType) {
		return -1
	}
	return list.ElemType[index]
}

func main() {
	//初始化
	list := SeqList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println(list.GetIndex(0))
	fmt.Println(list.GetIndex(1))
	list.RemoveList(1)
	fmt.Print(list.GetIndex(1))
}
