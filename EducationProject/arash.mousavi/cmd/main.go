package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	if _, err := fmt.Fscan(in, &n, &l); err != nil {
		return
	}

	weights := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &weights[i])
	}

	capacities := make([]int64, l)
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &capacities[i])
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	var score int64

	for _, w := range weights {
		bestLevel := -1
		var bestCap int64 = -1

		for level := 0; level < l; level++ {
			if capacities[level] >= w {
				if capacities[level] > bestCap ||
					(capacities[level] == bestCap && level > bestLevel) {
					bestCap = capacities[level]
					bestLevel = level
				}
			}
		}

		if bestLevel == -1 {

			return
		}

		capacities[bestLevel] -= w
		levelIndex := int64(bestLevel + 1)
		score += w * levelIndex
	}

	fmt.Println(score)
}
