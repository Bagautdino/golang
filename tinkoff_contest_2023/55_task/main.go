package main

import (
	"fmt"
	"sort"
)

type DisjointSetUnion struct {
	parent []int
	size   []int
}

func NewDisjointSetUnion(n int) *DisjointSetUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DisjointSetUnion{parent, size}
}

func (dsu *DisjointSetUnion) find(x int) int {
	if dsu.parent[x] == x {
		return x
	}
	dsu.parent[x] = dsu.find(dsu.parent[x])
	return dsu.parent[x]
}

func (dsu *DisjointSetUnion) union(x int, y int) bool {
	rootX := dsu.find(x)
	rootY := dsu.find(y)
	if rootX != rootY {
		if dsu.size[rootX] < dsu.size[rootY] {
			rootX, rootY = rootY, rootX
		}
		dsu.parent[rootY] = rootX
		dsu.size[rootX] += dsu.size[rootY]
		return true
	}
	return false
}

func main() {
	var numNodes, numEdges int
	fmt.Scan(&numNodes, &numEdges)

	dsu := NewDisjointSetUnion(numNodes)
	edges := make([][3]int, numEdges)

	for i := 0; i < numEdges; i++ {
		var from, to, weight int
		fmt.Scan(&from, &to, &weight)
		edges[i] = [3]int{from - 1, to - 1, weight}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] > edges[j][2]
	})

	maxWeight := edges[0][2]

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if dsu.union(from, to) {
			maxWeight = weight - 1
		}
	}

	fmt.Println(maxWeight)
}
