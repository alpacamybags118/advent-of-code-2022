package main

import (
	helpers "advent-of-code-2022/src/helpers"
	"flag"
	"fmt"
)

func main() {
	useSample := flag.Bool("use-sample", false, "Provide flag to run solution with sample data rather than input data")
	dayToRun := flag.String("part", "1", "Provide which part to solve. Defaults to 1")

	flag.Parse()

	if *dayToRun == "1" {
		Part1(*useSample)
	} else {
		Part2(*useSample)
	}
}

func Part1(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	positions := ProcessDatastreams(input, 4)

	fmt.Print(positions)
}

func Part2(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	positions := ProcessDatastreams(input, 14)

	fmt.Print(positions)
}

func ProcessDatastreams(datastreams []string, packetSize int) []int {
	var positions []int = make([]int, 0)

	for _, datastream := range datastreams {
		position := GetPositionOfStartOfPacketMarker(datastream, packetSize)

		positions = append(positions, position)
	}

	return positions
}

func GetPositionOfStartOfPacketMarker(datastream string, packetSize int) int {
	for i := packetSize; i < len(datastream); i++ {
		unique := true
		packet := datastream[i-packetSize : i]
		charsEncountered := make(map[rune]int)
		for _, char := range packet {
			if charsEncountered[char] != 0 {
				unique = false
				break
			}

			charsEncountered[char]++
		}

		if unique {
			return i
		}
	}

	return -1
}
