package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var parent []int
var rank []int
var groupSize []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nm := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	parent = make([]int, n+1)
	rank = make([]int, n+1)
	groupSize = make([]int, n+1)

	for i := 0; i <= n; i++ {
		parent[i] = i
	}

	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < m; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		command, _ := strconv.Atoi(input[0])

		if command == 1 {
			u, _ := strconv.Atoi(input[1])
			v, _ := strconv.Atoi(input[2])
			unionSets(u, v)
		} else if command == 2 {
			u, _ := strconv.Atoi(input[1])
			v, _ := strconv.Atoi(input[2])
			if areInSameSet(u, v) {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		} else if command == 3 {
			v, _ := strconv.Atoi(input[1])
			fmt.Fprintln(writer, setSize(v)+1)
		}
	}

	writer.Flush()
}

func findSet(v int) int {
	if parent[v] != v {
		parent[v] = findSet(parent[v])
	}
	return parent[v]
}

func areInSameSet(u, v int) bool {
	return findSet(u) == findSet(v)
}

func unionSets(u, v int) {
	rootU, rootV := findSet(u), findSet(v)
	if rootU == rootV {
		return
	}

	for vert := 1; vert < len(groupSize); vert++ {
		if areInSameSet(vert, rootU) || areInSameSet(vert, rootV) {
			groupSize[vert]++
		}
	}

	if rank[rootU] == rank[rootV] {
		rank[rootU]++
		parent[rootV] = rootU
	} else if rank[rootU] < rank[rootV] {
		parent[rootU] = rootV
	} else {
		parent[rootV] = rootU
	}
}

func setSize(v int) int {
	return groupSize[v]
}
