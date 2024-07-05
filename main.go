package main

import "fmt"

func main() {
	n := 10
	s := fibo(n)
	fmt.Println("fibo", s)
}

func fibo(n int) []int {
	fiboseri := make([]int, n)
	fiboseri[0], fiboseri[1] = 0, 1
	for i := 2; i < n; i++ {
		fiboseri[i] = fiboseri[i-1] + fiboseri[i-2]
	}
	return fiboseri
}
