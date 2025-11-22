package service

import (
	"amin_niazi/src/internal/model"
	"sort"
)

func ShelfAssigner(weights []int, capacities []int) ([]model.Shelf, []model.Item, int) {
	items := make([]model.Item, len(weights))
	for i, weight := range weights {
		items[i] = model.Item{
			Id:     i + 1,
			Weight: weight,
		}
	}

	L := len(capacities)

	storage := make([]model.Shelf, L)
	for i := 0; i < L; i++ {
		storage[i] = model.Shelf{
			Level:  i + 1,
			Filled: 0,
		}
	}
	totalCost := 0

	for li, capacity := range capacities {
		level := li + 1
		if len(items) == 0 {
			continue
		}
		thisLevelItems := knapsackSolver(items, capacity)
		if thisLevelItems == nil || len(thisLevelItems) == 0 {
			continue
		}

		dropMap := make(map[int]bool)
		var shelfItems []model.Item
		levelWeight := 0

		for _, itemIndex := range thisLevelItems {
			dropMap[itemIndex] = true
			item := items[itemIndex]
			shelfItems = append(shelfItems, item)
			levelWeight += item.Weight
		}
		totalCost += levelWeight * level
		storage[li].Items = shelfItems
		storage[li].Filled = levelWeight

		var leftOverItems []model.Item
		for itemIndex, item := range items {
			if !dropMap[itemIndex] {
				leftOverItems = append(leftOverItems, item)
			}
		}
		items = leftOverItems
	}

	// try to shove what's left somewhere
	leftOverItems := items
	if len(leftOverItems) > 0 {
		sort.Slice(leftOverItems, func(i, j int) bool {
			return leftOverItems[i].Weight > leftOverItems[j].Weight
		})

		var unlovedItems []model.Item
		for _, item := range leftOverItems {
			placed := false
			for lindex := 0; lindex < L; lindex++ {
				lcap := capacities[lindex]
				storageLevel := storage[lindex]
				used := storageLevel.Filled

				if used+item.Weight < lcap {
					storageLevel.Items = append(storageLevel.Items, item)
					storageLevel.Filled += item.Weight
					totalCost += item.Weight * storageLevel.Level
					placed = true
					break
				}
			}
			if !placed {
				unlovedItems = append(unlovedItems, item)
			}
		}
		leftOverItems = unlovedItems
	}

	return storage, leftOverItems, totalCost
}
