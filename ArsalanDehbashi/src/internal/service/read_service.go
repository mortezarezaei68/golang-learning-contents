package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ReadService interface {
	ReadData() ([]int, []int)
}

type readService struct{}

func NewReadService() ReadService {
	return &readService{}
}

func (s *readService) ReadData() ([]int, []int) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter N and L:")
	if !reader.Scan() {
		log.Fatal("failed to read first line")
	}
	first := strings.Fields(reader.Text())
	if len(first) != 2 {
		log.Fatal("expected 2 numbers")
	}

	N, err := strconv.Atoi(first[0])
	if err != nil {
		log.Fatal("invalid N:", err)
	}

	L, err := strconv.Atoi(first[1])
	if err != nil {
		log.Fatal("invalid L:", err)
	}

	fmt.Println("Enter weights of products:")
	if !reader.Scan() {
		log.Fatal("failed to read second line")
	}
	weighFields := strings.Fields(reader.Text())
	if len(weighFields) != N {
		log.Fatalf("expected %d weights", N)
	}

	productWeights := make([]int, N)
	for i, f := range weighFields {
		productWeights[i], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal("invalid weight:", err)
		}
	}

	fmt.Println("Enter capacities of shelves:")
	if !reader.Scan() {
		log.Fatal("failed to read third line")
	}
	capFields := strings.Fields(reader.Text())
	if len(capFields) != L {
		log.Fatalf("expected %d capacities", L)
	}

	levelCapacities := make([]int, L)
	for i, f := range capFields {
		levelCapacities[i], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal("invalid capacity:", err)
		}
	}

	return productWeights, levelCapacities
}
