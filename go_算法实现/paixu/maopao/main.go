package main

import "fmt"

func main() {
	var arr = []int{5, 6, 7, 1, 23}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j] > arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Print(arr)
}
