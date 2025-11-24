package main

import (
	"fmt"
	"sort"
)

func Max(n, l int, weight []int, capacities []int) int64 {
	sort.Ints(weight)
	type Shelf struct {
		level    int
		capacity int
	}
	shelevs := make([]Shelf, l)
	for i := 0; i < l; i++ {
		shelevs[i] = Shelf{level: i + 1, capacity: capacities[i]}
	}
	var totalScore int64 = 0
	productIndex := 0

	for _, shelf := range shelevs {
		currentLevel := shelf.level
		remainingCapacity := shelf.capacity

		for remainingCapacity > 0 && productIndex < n {
			productWeight := weight[productIndex]

			if productWeight <= remainingCapacity {
				totalScore += int64(productWeight) * int64(currentLevel)
				remainingCapacity -= productWeight
				productIndex++
			} else {
				break
			}
		}
	}

	return totalScore
}

func main() {

	N := 4
	L := 2
	weights := []int{3, 2, 5, 4}
	capacities := []int{7, 10}

	maxScore := Max(N, L, weights, capacities)

	fmt.Printf(" maximum capacity :) %d\n", maxScore)
}
