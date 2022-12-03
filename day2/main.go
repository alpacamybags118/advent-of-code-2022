package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Move struct {
	opponentSymbol string
	playerSymbol   string
	value          int
}

var ROCK Move = Move{
	"A",
	"X",
	1,
}

var PAPER Move = Move{
	"B",
	"Y",
	2,
}

var SCISSORS Move = Move{
	"C",
	"Z",
	3,
}

func main() {
	//Part1()
	Part2()
}

func Part1() {
	games, err := ReadInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	score := CalculateScore(games, true)

	fmt.Println(score)
}

func Part2() {
	games, err := ReadInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	score := CalculateScore(games, false)

	fmt.Println(score)
}

func ReadInput() ([]string, error) {
	var matches []string = make([]string, 0)

	readFile, err := os.Open("input")

	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		matches = append(matches, scanner.Text())
	}

	return matches, nil
}

func CalculateScore(matches []string, isStandard bool) int {
	var totalScore int = 0

	for i, match := range matches {
		var matchScore int = 0

		plays := GetPlays(match)
		if i == 0 {
			fmt.Println(plays)
		}

		if isStandard {
			matchScore += plays[1].value
			matchScore += GetGameScore(plays[0], plays[1])
		} else {
			matchScore += GetGoalMatchScore(plays[0], plays[1], i)
		}

		fmt.Println(matchScore)

		totalScore += matchScore
	}

	return totalScore
}

func GetPlays(match string) []Move {
	var moves []Move = make([]Move, 0)

	plays := strings.Split(match, " ")

	for _, play := range plays {
		switch play {
		case ROCK.opponentSymbol, ROCK.playerSymbol:
			moves = append(moves, ROCK)
		case PAPER.opponentSymbol, PAPER.playerSymbol:
			moves = append(moves, PAPER)
		case SCISSORS.opponentSymbol, SCISSORS.playerSymbol:
			moves = append(moves, SCISSORS)
		}
	}

	return moves
}

func GetGameScore(opponent Move, player Move) int {
	fmt.Println(opponent.opponentSymbol == player.opponentSymbol)
	if opponent.opponentSymbol == player.opponentSymbol {
		fmt.Println("draw")
		return 3
	}

	switch opponent.opponentSymbol {
	case ROCK.opponentSymbol:
		if player.playerSymbol == PAPER.playerSymbol {
			fmt.Println("win")
			return 6
		}
	case PAPER.opponentSymbol:
		if player.playerSymbol == SCISSORS.playerSymbol {
			fmt.Println("win")
			return 6
		}
	case SCISSORS.opponentSymbol:
		if player.playerSymbol == ROCK.playerSymbol {
			fmt.Println("win")
			return 6
		}
	}

	fmt.Println("loss")
	return 0
}

func GetGoalMatchScore(opponent Move, player Move, index int) int {
	var score int = 0
	fmt.Printf("score: %v \n", GetPlayerScoreForGoalMatch(opponent, "draw"))
	switch player.opponentSymbol {
	case ROCK.opponentSymbol: //loss
		fmt.Println("loss")
		score += GetPlayerScoreForGoalMatch(opponent, "loss") + 0
	case PAPER.opponentSymbol: //draw
		fmt.Println("draw")
		score += GetPlayerScoreForGoalMatch(opponent, "draw") + 3
	case SCISSORS.opponentSymbol: //win
		fmt.Println("win")
		score += GetPlayerScoreForGoalMatch(opponent, "win") + 6
	}

	return score
}

func GetPlayerScoreForGoalMatch(opponent Move, gameType string) int {
	if gameType == "draw" {
		return opponent.value
	} else if gameType == "loss" {
		switch opponent.opponentSymbol {
		case ROCK.opponentSymbol: //loss
			return SCISSORS.value
		case PAPER.opponentSymbol: //draw
			return ROCK.value
		case SCISSORS.opponentSymbol: //win
			return PAPER.value
		}
	}

	switch opponent.opponentSymbol {
	case ROCK.opponentSymbol: //loss
		return PAPER.value
	case PAPER.opponentSymbol: //draw
		return SCISSORS.value
	case SCISSORS.opponentSymbol: //win
		return ROCK.value
	}

	return 0
}
