package main

import (
	helpers "advent-of-code-2022/src/lib"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	NumberToMove int
	FromStack    int
	ToStack      int
}

func main() {
	useSample := flag.Bool("use-sample", false, "Provide flag to run solution with sample data rather than input data")
	dayToRun := flag.String("part", "1", "Provide which part to solve. Defaults to 1")

	flag.Parse()
	fmt.Println(*useSample)
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

	out := RunCraneSimuation(input, false)

	for _, queue := range out {
		fmt.Print(queue[len(queue)-1])
	}
}

func Part2(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	out := RunCraneSimuation(input, true)

	for _, queue := range out {
		fmt.Print(queue[len(queue)-1])
	}
}

func RunCraneSimuation(input []string, moveAll bool) [][]string {
	cratePos := GetInitialCratePositions(input)
	instructions := GetInstructions(input)

	finalCratePos := RunInstructions(cratePos, instructions, moveAll)

	return finalCratePos
}

func GetInitialCratePositions(input []string) [][]string {
	var sepIndex int = 0

	for index, item := range input {
		if item == "" {
			sepIndex = index
			break
		}
	}
	crateMap := input[0 : sepIndex-1]
	numQueues := len(strings.ReplaceAll(input[sepIndex-1], " ", ""))

	var crateQueues [][]string = make([][]string, numQueues)

	for i := len(crateMap) - 1; i >= 0; i-- {
		// starting index is 1, then +4 each time until end
		crateQueueIndex := 0
		startingIndex := 1

		for x := startingIndex; x < len(crateMap[i]); x = x + 4 {
			if len(crateQueues) < crateQueueIndex+1 {
				crateQueues[crateQueueIndex] = make([]string, 0)
			}

			if string(crateMap[i][x]) != " " {
				crateQueues[crateQueueIndex] = append(crateQueues[crateQueueIndex], string(crateMap[i][x]))
			}

			crateQueueIndex += 1
		}
	}

	return crateQueues
}

func GetInstructions(input []string) []Instruction {
	var output []Instruction = make([]Instruction, 0)
	var sepIndex int = 0

	for index, item := range input {
		if item == "" {
			sepIndex = index
			break
		}
	}

	instructions := input[sepIndex+1:]

	for _, instructionString := range instructions {
		instructionString = strings.Replace(instructionString, "move", "", -1)
		instructionString = strings.Replace(instructionString, "from", "", -1)
		instructionString = strings.Replace(instructionString, "to", "", -1)

		instructionList := strings.Split(instructionString, "  ")

		numMove, _ := strconv.Atoi(strings.Trim(instructionList[0], " "))
		from, _ := strconv.Atoi(strings.Trim(instructionList[1], " "))
		to, _ := strconv.Atoi(strings.Trim(instructionList[2], " "))

		output = append(output, Instruction{
			numMove,
			from - 1,
			to - 1,
		})
	}

	return output
}

func RunInstructions(crateQueues [][]string, instructions []Instruction, moveAll bool) [][]string {
	for _, instruction := range instructions {
		if moveAll {
			tmp := crateQueues[instruction.FromStack][len(crateQueues[instruction.FromStack])-instruction.NumberToMove:]

			crateQueues[instruction.FromStack] = crateQueues[instruction.FromStack][0 : len(crateQueues[instruction.FromStack])-instruction.NumberToMove]
			for i := 0; i < len(tmp); i++ {
				crateQueues[instruction.ToStack] = append(crateQueues[instruction.ToStack], tmp[i])
			}

		} else {
			for i := 0; i < instruction.NumberToMove; i++ {
				tmp := crateQueues[instruction.FromStack][len(crateQueues[instruction.FromStack])-1]
				crateQueues[instruction.FromStack] = crateQueues[instruction.FromStack][0 : len(crateQueues[instruction.FromStack])-1]
				crateQueues[instruction.ToStack] = append(crateQueues[instruction.ToStack], tmp)
			}
		}

	}

	return crateQueues
}
