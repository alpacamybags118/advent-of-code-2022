package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var useSample bool = false

	if len(os.Args) >= 2 {
		useSample = true
	}

	//Part1(useSample)
	Part2(useSample)
}

func Part1(useSample bool) {
	input, err := ReadInput(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	numPairs := GetNumPairs(input)

	fmt.Println(numPairs)
}

func Part2(useSample bool) {
	input, err := ReadInput(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	numPairs := GetNumPairsWithAnyOverlap(input)

	fmt.Println(numPairs)
}

func ReadInput(useSample bool) ([]string, error) {
	var ranges []string = make([]string, 0)
	var file string = "input"

	if useSample {
		file = "sample"
	}

	readFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ranges = append(ranges, scanner.Text())
	}

	return ranges, nil
}

func GetNumPairs(input []string) int {
	var totalPairs int = 0

	for _, pair := range input {
		var pairA []int = make([]int, 0)
		var pairB []int = make([]int, 0)

		splitPair := strings.Split(pair, ",")
		fmt.Println(splitPair)

		for _, item := range strings.Split(splitPair[0], "-") {
			num, _ := strconv.Atoi(item)
			pairA = append(pairA, num)
		}

		for _, item := range strings.Split(splitPair[1], "-") {
			num, _ := strconv.Atoi(item)
			pairB = append(pairB, num)
		}

		if IsFullOverlap(pairA, pairB) {
			totalPairs++
		}

	}

	return totalPairs
}

func GetNumPairsWithAnyOverlap(input []string) int {
	var totalPairs int = 0

	for _, pair := range input {
		var pairA []int = make([]int, 0)
		var pairB []int = make([]int, 0)

		splitPair := strings.Split(pair, ",")
		fmt.Println(splitPair)

		for _, item := range strings.Split(splitPair[0], "-") {
			num, _ := strconv.Atoi(item)
			pairA = append(pairA, num)
		}

		for _, item := range strings.Split(splitPair[1], "-") {
			num, _ := strconv.Atoi(item)
			pairB = append(pairB, num)
		}

		if IsAnyOverlap(pairA, pairB) {
			totalPairs++
		}

	}

	return totalPairs
}

func IsFullOverlap(pair1 []int, pair2 []int) bool {
	fmt.Println(pair1)
	fmt.Println(pair2)
	fmt.Println((pair1[0]-pair2[0] <= 0 && pair1[1]-pair2[1] >= 0) || (pair2[0]-pair1[0] <= 0 && pair2[1]-pair1[1] >= 0))
	return (pair1[0]-pair2[0] <= 0 && pair1[1]-pair2[1] >= 0) || (pair2[0]-pair1[0] <= 0 && pair2[1]-pair1[1] >= 0)
}

func IsAnyOverlap(pair1 []int, pair2 []int) bool {
	var pairARange []int = make([]int, 0)
	var pairBRange []int = make([]int, 0)

	for i := pair1[0]; i <= pair1[1]; i++ {
		pairARange = append(pairARange, i)
	}

	for i := pair2[0]; i <= pair2[1]; i++ {
		pairBRange = append(pairBRange, i)
	}

	for _, numA := range pairARange {
		for _, numB := range pairBRange {
			if numB == numA {
				return true
			}
		}
	}

	return false
}
