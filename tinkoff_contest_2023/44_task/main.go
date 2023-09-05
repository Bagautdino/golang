package main

import (
	"fmt"
	"sort"
)

func main() {
	var targetSum, numCoins int
	fmt.Scan(&targetSum, &numCoins)

	coinValues := make([]int, numCoins)
	for i := 0; i < numCoins; i++ {
		fmt.Scan(&coinValues[i])
	}

	sort.Ints(coinValues)

	const maxCount = 2

	result := make([]int, 0, 2*numCoins)
	count := 0
	sum := 0

	for i := numCoins - 1; i >= 0; i-- {
		if sum >= targetSum {
			break
		}

		curValue := coinValues[i]
		needCount := (targetSum - sum) / curValue

		if needCount > maxCount {
			needCount = maxCount
		}

		for j := 0; j < needCount; j++ {
			result = append(result, curValue)
			count++
			sum += curValue
		}
	}

	if sum != targetSum {
		count = -1
	}

	fmt.Println(count)

	if count != -1 {
		for i := 0; i < count; i++ {
			fmt.Print(result[count-i-1], " ")
		}
	}
}
