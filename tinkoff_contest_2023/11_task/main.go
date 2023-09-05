package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	max := 0
	for i := 0; i < n; i++ {
		if a[i] < s && a[i] > max {
			max = a[i]
		}

	}
	fmt.Println(max)
}
