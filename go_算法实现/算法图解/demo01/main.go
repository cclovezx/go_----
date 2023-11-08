package main

import "fmt"

//递归99乘法表
func main() {
	multiplication(1)
}

func multiplication(n int) {
	if n < 10 {
		for m := 1; m < n+1; m++ {
			fmt.Printf("%d * %d = %d\t", m, n, m*n)
		}
		fmt.Println()
		multiplication(n + 1)
	}
}
