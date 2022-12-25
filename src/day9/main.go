package main

import (
	helpers "advent-of-code-2022/src/helpers"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Grid struct {
	head Coordinate
	tail Coordinate
}

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
	var rope []Coordinate = []Coordinate{{x: 0, y: 0}, {x: 0, y: 0}}

	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}
	TraverseGrid(input, rope)
}

func Part2(useSample bool) {
	var rope []Coordinate = []Coordinate{{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}}
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	TraverseGrid(input, rope)
}

func TraverseGrid(input []string, rope []Coordinate) {
	var visited map[string]int = make(map[string]int)

	visited["0,0"] = 1

	for _, step := range input {
		var direction string = strings.Split(step, " ")[0]
		var distance int64

		distance, _ = strconv.ParseInt(strings.Split(step, " ")[1], 10, 64)

		rope, visited = UpdatePositions(rope, direction, distance, visited)
	}

	fmt.Println(len(visited))
}

func UpdatePositions(rope []Coordinate, direction string, distance int64, visited map[string]int) ([]Coordinate, map[string]int) {
	fmt.Println("starting")
	PrintKnotPositions(rope)
	fmt.Println()
	for distance > 0 {
		switch direction {
		case "R":
			rope[0].x++
			for i := 1; i < len(rope); i++ {
				if !AreTouching(rope[i-1], rope[i]) {
					if rope[i].x == rope[i-1].x {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}
					} else if MustMoveDiagonally(rope[i-1], rope[i]) {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}

						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else {
						rope[i].x++
					}
				}
			}
			break
		case "U":
			rope[0].y++
			for i := 1; i < len(rope); i++ {
				//fmt.Println(AreTouching(rope[i-1], rope[i]))
				if !AreTouching(rope[i-1], rope[i]) {
					if rope[i].y == rope[i-1].y {
						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else if MustMoveDiagonally(rope[i-1], rope[i]) {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}

						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else {
						rope[i].y++
					}
				}
			}
			break
		case "L":
			rope[0].x--
			for i := 1; i < len(rope); i++ {
				if !AreTouching(rope[i-1], rope[i]) {
					if rope[i].x == rope[i-1].x {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}
					} else if MustMoveDiagonally(rope[i-1], rope[i]) {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}

						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else {
						rope[i].x--
					}
				}
			}
			fmt.Println("positions")
			PrintKnotPositions(rope)
			fmt.Println()
			break
		case "D":
			rope[0].y--
			for i := 1; i < len(rope); i++ {
				//fmt.Println(AreTouching(rope[i-1], rope[i]))
				if !AreTouching(rope[i-1], rope[i]) {
					if rope[i].y == rope[i-1].y {
						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else if MustMoveDiagonally(rope[i-1], rope[i]) {
						if rope[i].y < rope[i-1].y {
							rope[i].y++
						} else if rope[i].y > rope[i-1].y {
							rope[i].y--
						}

						if rope[i].x < rope[i-1].x {
							rope[i].x++
						} else if rope[i].x > rope[i-1].x {
							rope[i].x--
						}
					} else {
						rope[i].y--
					}
				}
			}
			break
		}
		visited[fmt.Sprintf("%v,%v", rope[len(rope)-1].x, rope[len(rope)-1].y)] = 1

		distance--

	}
	//PrintKnotPositions(rope)

	return rope, visited
}

func AreTouching(head Coordinate, tail Coordinate) bool {
	var xDiff = math.Abs(float64(head.x - tail.x))
	var yDiff = math.Abs(float64(head.y - tail.y))

	return xDiff <= 1 && yDiff <= 1
}

func MustMoveDiagonally(head Coordinate, tail Coordinate) bool {
	var xDiff = math.Abs(float64(head.x - tail.x))
	var yDiff = math.Abs(float64(head.y - tail.y))

	return xDiff > 1 || yDiff > 1
}

func PrintKnotPositions(rope []Coordinate) {
	for index, knot := range rope {
		fmt.Printf("Knot %v: %v,%v\n", index, knot.x, knot.y)
	}

	fmt.Println()
}

/*
1,0 0,0 - touching
2,0 0,0 - not touching
*/
