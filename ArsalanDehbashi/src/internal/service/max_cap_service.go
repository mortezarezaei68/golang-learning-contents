package service

import (
	"fmt"
	"maxshelfstabilityscore/src/internal/model"
	"sort"
)

type MaxCapService interface {
	CalculateMaxCap([]model.Product, []model.Level) (int, error)
	NameProducts([]int) ([]model.Product, error)
	NameLevels([]int) ([]model.Level, error)
}

type maxCapService struct{}

func NewMaxCapService() MaxCapService {
	return &maxCapService{}
}

func (s *maxCapService) CalculateMaxCap(products []model.Product, levels []model.Level) (int, error) {
	sortProductsDesc(products)
	sortLevelsDesc(levels)

	for i := range products {
		placed := false
		for j := range levels {
			if levels[j].RemainingCap >= products[i].Weight {
				if j == len(levels)-1 || levels[j].RemainingCap >= levels[j+1].RemainingCap {
					products[i].Level = levels[j].Order
					levels[j].RemainingCap -= products[i].Weight
					placed = true
					break
				}

				continue
			}
		}

		if !placed {
			return 0, fmt.Errorf("there isn't enough capacity to place all the products")
		}
	}

	fmt.Println(products)

	score := 0
	for _, p := range products {
		score += p.Weight * p.Level
	}

	return score, nil
}

func (s *maxCapService) NameProducts(productWeights []int) ([]model.Product, error) {
	products := make([]model.Product, len(productWeights))
	for i := range productWeights {
		products[i].Name = fmt.Sprintf("P%d", i+1)
		products[i].Weight = productWeights[i]
	}
	return products, nil
}

func (s *maxCapService) NameLevels(levelCapacities []int) ([]model.Level, error) {
	levels := make([]model.Level, len(levelCapacities))
	for i := range levelCapacities {
		levels[i].Order = i + 1
		levels[i].Capacity = levelCapacities[i]
		levels[i].RemainingCap = levelCapacities[i]
	}
	return levels, nil
}

func sortProductsDesc(products []model.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Weight > products[j].Weight
	})
}

func sortLevelsDesc(levels []model.Level) {
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].Capacity > levels[j].Capacity
	})
}
