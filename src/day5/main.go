package main

import (
	helpers "advent-of-code-2022/src/lib"
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

	fmt.Println(len(input))
}

func Part2(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(input)
}
