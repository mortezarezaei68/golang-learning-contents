package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	var NumberOfProducts, NumberOfShelves, MaximumAchievableStabilityScore int

	type Product struct {
		Weight       int
		LevelOfShelf int
	}

	type Shelf struct {
		Weight       int
		RemainWeight int
	}

	var Products []Product
	var Shelves []Shelf

	labels := map[int]string{
		0: "number of Products : ",
		1: "number of Shelves : ",
		2: "weight of Products : ",
		3: "max weight capacity of Shelves : ",
	}

	for i := 0; i < 4; i++ {
		fmt.Print(labels[i])

		if !scanner.Scan() {
			return
		}

		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if i == 0 {
			v, _ := strconv.Atoi(line)
			NumberOfProducts = v
		} else if i == 1 {
			v, _ := strconv.Atoi(line)
			NumberOfShelves = v
		} else if i == 2 {
			for _, value := range strings.Fields(line) {
				weight, _ := strconv.Atoi(value)
				Products = append(Products, Product{weight, 0})
			}
		} else if i == 3 {
			for _, value := range strings.Fields(line) {
				weight, _ := strconv.Atoi(value)
				Shelves = append(Shelves, Shelf{weight, weight})
			}
		}
	}

	sort.Slice(Products, func(i, j int) bool {
		return Products[i].Weight < Products[j].Weight
	})

	sort.Slice(Shelves, func(i, j int) bool {
		return Shelves[i].Weight < Shelves[j].Weight
	})

	for i := 0; i < NumberOfShelves; i++ {
		for j := 0; j < NumberOfProducts; j++ {
			if Products[j].LevelOfShelf != 0 {
				continue
			}
			if Shelves[i].RemainWeight >= Products[j].Weight {
				Shelves[i].RemainWeight -= Products[j].Weight
				Products[j].LevelOfShelf = i + 1
			}
		}
	}

	for _, product := range Products {
		MaximumAchievableStabilityScore += product.Weight * product.LevelOfShelf
	}

	fmt.Println("Maximum achievable stability score : ", MaximumAchievableStabilityScore)
}
