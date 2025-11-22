package main

import (
	"fmt"
	"log"
	"maxshelfstabilityscore/src/internal/service"
)

func main() {
	readSvs := service.NewReadService()
	maxCapSvc := service.NewMaxCapService()

	productWeights, levelCapacities := readSvs.ReadData()

	products, _ := maxCapSvc.NameProducts(productWeights)
	levels, _ := maxCapSvc.NameLevels(levelCapacities)

	maxScore, err := maxCapSvc.CalculateMaxCap(products, levels)
	if err != nil {
		log.Fatal("the error is: ", err)
	}
	fmt.Println(maxScore)
}
