package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInputAsStringArray(useSample bool) ([]string, error) {
	var input []string = make([]string, 0)
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
		input = append(input, scanner.Text())
	}

	return input, nil
}

func ReadInputAsIntGrid(useSample bool) ([][]int, error) {
	var grid [][]int = make([][]int, 0)

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
		var row []int = make([]int, 0)
		rowText := strings.Split(scanner.Text(), "")

		for _, num := range rowText {
			rowNum, err := strconv.Atoi(num)

			if err != nil {
				return nil, err
			}

			row = append(row, rowNum)
		}

		grid = append(grid, row)
	}

	return grid, nil
}
