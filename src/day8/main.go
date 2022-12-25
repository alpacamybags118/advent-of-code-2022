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
	input, err := helpers.ReadInputAsIntGrid(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	totalVisibleTrees := FindVisibleTrees(input)

	fmt.Println(totalVisibleTrees)
}

func Part2(useSample bool) {
	input, err := helpers.ReadInputAsIntGrid(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	largestScenicScore := GetLargestScenicScore(input)

	fmt.Println(largestScenicScore)
}

func FindVisibleTrees(grid [][]int) int {
	var treesVisible int = 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			if column == 0 || column == len(grid)-1 || row == 0 || row == len(grid)-1 { //outer grid case
				//fmt.Printf("%v,%v is visible\n", row, column)
				treesVisible++
				continue
			}

			if IsVisible(row, column, grid) {
				treesVisible++
			}
		}
	}

	return treesVisible
}

func GetLargestScenicScore(grid [][]int) int {
	var largestScenicStore int = 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			if column == 0 || column == len(grid)-1 || row == 0 || row == len(grid)-1 { //outer grid case
				continue
			}

			scenicScore := GetScenicScore(row, column, grid)

			if scenicScore > largestScenicStore {
				largestScenicStore = scenicScore
			}
		}
	}

	return largestScenicStore
}

func IsVisible(row int, column int, grid [][]int) bool {
	var up int = row
	var down int = row
	var right int = column
	var left int = column
	var treeSize int = grid[row][column]
	var upVisible bool = true
	var downVisible bool = true
	var leftVisible bool = true
	var rightVisible bool = true

	for up >= 0 {
		up--
		if up >= 0 {
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[up][column])
			if grid[up][column] >= treeSize {
				upVisible = false
				break
			}
		}
	}

	for down < len(grid) {
		down++

		if down < len(grid) {
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[down][column])
			if grid[down][column] >= treeSize {
				downVisible = false
				break
			}
		}
	}

	for left >= 0 {
		left--

		if left >= 0 {
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[row][left])
			if grid[row][left] >= treeSize {
				leftVisible = false
				break
			}
		}
	}

	for right < len(grid) {
		right++

		if right < len(grid) {
			if grid[row][right] >= treeSize {
				rightVisible = false
				break
			}
		}
	}

	if upVisible || downVisible || leftVisible || rightVisible {
		fmt.Printf("%v,%v is visible\n", row, column)
	}

	return upVisible || downVisible || leftVisible || rightVisible
}

func GetScenicScore(row int, column int, grid [][]int) int {
	var up int = row
	var down int = row
	var right int = column
	var left int = column

	var treeSize int = grid[row][column]

	var upVisible int = 0
	var downVisible int = 0
	var leftVisible int = 0
	var rightVisible int = 0

	for up >= 0 {
		up--
		if up >= 0 {
			upVisible++
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[up][column])
			if grid[up][column] >= treeSize {
				break
			}
		}
	}

	for down < len(grid) {
		down++
		if down < len(grid) {
			downVisible++
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[down][column])
			if grid[down][column] >= treeSize {
				break
			}
		}
	}

	for left >= 0 {
		left--
		if left >= 0 {
			leftVisible++
			fmt.Printf("comparing original tree %v and %v\n", treeSize, grid[row][left])
			if grid[row][left] >= treeSize {
				break
			}
		}
	}

	for right < len(grid) {
		right++

		if right < len(grid) {
			rightVisible++
			if grid[row][right] >= treeSize {
				break
			}
		}
	}

	return upVisible * downVisible * leftVisible * rightVisible
}
