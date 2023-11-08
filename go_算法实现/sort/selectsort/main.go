package main

import "fmt"

func main() {
	var a = []int{6, 1, 3, 34, 54, 7, 8, 9, 10}
	for i := 0; i < len(a); i++ {
		point := i
		for j := i + 1; j < len(a); j++ {
			if a[point] > a[j] {
				point = j
			}
		}
		temp := a[i]
		a[i] = a[point]
		a[point] = temp
	}
	fmt.Println(a)
}
