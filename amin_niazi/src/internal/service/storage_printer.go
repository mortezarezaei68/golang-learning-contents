package service

import (
	"amin_niazi/src/internal/model"
	"fmt"
	"strings"
)

func StoragePrinterService(storage []model.Shelf, leftovers []model.Item, totalCost int) {
	for _, level := range storage {
		fmt.Println("-------------------------------")
		fmt.Printf("Level: %d - filled: %d\n", level.Level, level.Filled)
		fmt.Println("------------ items ------------")
		var toPrint []string
		for _, item := range level.Items {
			toPrint = append(toPrint, fmt.Sprintf("P%d (w=%d)", item.Id, item.Weight))
		}
		fmt.Println(strings.Join(toPrint, ", "))
	}
	fmt.Println("------------------------------")
	fmt.Printf("Total Cost: %d\n", totalCost)

	if len(leftovers) > 0 {
		fmt.Println("-------------- leftovers ------------")
		for _, leftover := range leftovers {
			fmt.Printf("P%d\n - W%d", leftover.Id, leftover.Weight)
		}
	}
}
