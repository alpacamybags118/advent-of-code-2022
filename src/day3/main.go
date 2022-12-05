package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Rucksack struct {
	compartmentA map[string]int
	compartmentB map[string]int
}

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

	sacks := GetRucksacks(input)

	var matchTotal int = 0
	for _, sack := range sacks {
		for item := range sack.compartmentA {
			if _, exists := sack.compartmentB[item]; exists {
				matchTotal += strings.Index(alphabet, item) + 1
			}
		}
	}

	fmt.Println(matchTotal)
}

func Part2(useSample bool) {
	input, err := ReadInput(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	inventorys := GetGroupInventory(input)

	var inventoryTotal int = 0
	for _, inventory := range inventorys {
		for item, total := range inventory {
			if total == 3 {
				inventoryTotal += strings.Index(alphabet, item) + 1
			}
		}
	}

	fmt.Println(inventoryTotal)
}

func ReadInput(useSample bool) ([]string, error) {
	var sacks []string = make([]string, 0)
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
		sacks = append(sacks, scanner.Text())
	}

	return sacks, nil
}

func GetRucksacks(inputs []string) []Rucksack {
	var ruckSacks []Rucksack = make([]Rucksack, 0)

	for _, input := range inputs {
		ruckSack := Rucksack{
			make(map[string]int),
			make(map[string]int),
		}

		sackA := []rune(input)[0 : len(input)/2]
		sackB := []rune(input)[len(input)/2 : len(input)]

		for _, item := range sackA {
			key := string(item)
			ruckSack.compartmentA[key] += 1
		}

		for _, item := range sackB {
			key := string(item)
			ruckSack.compartmentB[key] += 1
		}

		ruckSacks = append(ruckSacks, ruckSack)
	}

	return ruckSacks
}

func GetGroupInventory(inputs []string) []map[string]int {
	var groupsInventory []map[string]int = make([]map[string]int, 0)

	for i := 0; i < len(inputs); i += 3 {
		var groupInventory map[string]int = make(map[string]int)
		for j := i; j < i+3; j++ {
			uniqueInventory := RemoveDupes(inputs[j])

			for _, item := range uniqueInventory {
				groupInventory[string(item)] += 1
			}
		}
		groupsInventory = append(groupsInventory, groupInventory)
	}

	return groupsInventory
}

func RemoveDupes(input string) string {
	var output string = ""

	for _, item := range input {
		if !strings.Contains(output, string(item)) {
			output = fmt.Sprintf("%s%s", output, string(item))
		}
	}

	return output
}
