package helpers

import (
	"bufio"
	"os"
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
