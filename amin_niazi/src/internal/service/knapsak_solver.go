package service

import "amin_niazi/src/internal/model"

func knapsackSolver(items []model.Item, capacity int) []int {
	n := len(items)
	if n == 0 || capacity <= 0 {
		return nil
	}
	dp := make([]int, capacity+1)
	choose := make([][]bool, n+1)
	for i := range choose {
		choose[i] = make([]bool, capacity+1)
	}
	for i := 0; i < n; i++ {
		currentItemWeight := items[i].Weight
		for j := capacity; j >= currentItemWeight; j-- {
			if dp[j-currentItemWeight]+currentItemWeight > dp[j] {
				dp[j] = dp[j-currentItemWeight] + currentItemWeight
				choose[i][j] = true
			}
		}
	}

	var result []int
	j := capacity
	for i := n - 1; i >= 0; i-- {
		if choose[i][j] {
			result = append(result, i)
			j -= items[i].Weight
		}
	}
	return result

}
