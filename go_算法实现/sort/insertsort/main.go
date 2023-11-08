// 选择排序
package main

import "fmt"

func main() {
	arr := []int{1, 23, 2, 2145, 2356, 4, 7, 3, 21, 23}
	for i := 0; i < len(arr); i++ {
		k := i
		for k > 0 {
			if arr[k] < arr[k-1] {
				temp := arr[k-1]
				arr[k-1] = arr[k]
				arr[k] = temp
			} else {
				break
			}
			k--
		}
	}
	fmt.Println(arr)
}
