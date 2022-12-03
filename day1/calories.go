package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Part1()
	Part2()
}

func Part1() {
	elfCollection, err := ReadInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	highestCalories := GetHighestCalories(elfCollection)

	fmt.Println(highestCalories)
}

func Part2() {
	elfCollection, err := ReadInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	topCalories := GetTopNCalories(elfCollection, 3)

	var sum int = 0

	for _, calories := range topCalories {
		sum += calories
	}

	fmt.Println(sum)
}

func ReadInput() ([][]int, error) {
	var elfCollection [][]int = make([][]int, 0)

	readFile, err := os.Open("input")

	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)

	var elfFoodCollection []int = make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfCollection = append(elfCollection, elfFoodCollection)
			elfFoodCollection = make([]int, 0)
		} else {
			calories, err := strconv.Atoi(line)

			if err != nil {
				return nil, err
			}

			elfFoodCollection = append(elfFoodCollection, calories)
		}
	}

	elfCollection = append(elfCollection, elfFoodCollection)

	return elfCollection, nil
}

func GetHighestCalories(elfCollection [][]int) int {
	var highestCalories int = 0

	for _, elfFood := range elfCollection {
		var totalCalories int = 0

		for _, calorie := range elfFood {
			totalCalories += calorie
		}

		if totalCalories > highestCalories {
			highestCalories = totalCalories
		}
	}

	return highestCalories
}

func GetTopNCalories(elfCollection [][]int, n int) []int {
	var topNCalories []int = make([]int, 0)

	for _, elfFood := range elfCollection {
		var totalCalories int = 0

		for _, calorie := range elfFood {
			totalCalories += calorie
		}

		topNCalories = AddToTopCalories(topNCalories, n, totalCalories)
	}

	return topNCalories
}

func AddToTopCalories(topCaloriesArray []int, n int, totalCalories int) []int {
	if len(topCaloriesArray) < n {
		return append(topCaloriesArray, totalCalories)
	}

	for i, topCalorie := range topCaloriesArray {
		if totalCalories > topCalorie {

			topCaloriesArray[i] = totalCalories

			return topCaloriesArray
		}
	}

	return topCaloriesArray
}
