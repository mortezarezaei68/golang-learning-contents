package main

import (
	"amin_niazi/src/internal/service"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, L int
	fmt.Fscan(reader, &N, &L)
	weights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &weights[i])
	}
	capacities := make([]int, L)
	for i := 0; i < L; i++ {
		fmt.Fscan(reader, &capacities[i])
	}
	storage, leftovers, totalCost := service.ShelfAssigner(weights, capacities)
	service.StoragePrinterService(storage, leftovers, totalCost)
}
