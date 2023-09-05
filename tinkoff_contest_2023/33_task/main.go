package main

import (
	"fmt"
	"sort"
)

func scanArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	scanArray(a)
	scanArray(b)
	l := 0
	r := n - 1
	for l < n && a[l] == b[l] {
		l++
	}
	for r >= 0 && a[r] == b[r] {
		r--
	}
	if l > r {
		fmt.Print("YES")
	} else {
		sort.Ints(a[l : r+1])
		if compareSlices(a[l:r+1], b[l:r+1]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func compareSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
