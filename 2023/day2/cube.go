package day2

// https://adventofcode.com/2023/day/2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// parseGames Parses the game input and returns a list of game objects
func parseGames(inputFile string) []Game {
	var parsedGames []Game
	var inputLines []string

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	for _, line := range inputLines {
		turnString := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(turnString[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		gameTurns := strings.Split(strings.TrimSpace(turnString[1]), ";")
		newGame := new(Game)
		parsedTurns := newGame.ParseTurns(gameTurns)
		newGame.New(id, parsedTurns)
		parsedGames = append(parsedGames, *newGame)
	}
	return parsedGames
}

// getSumOfValidGames Takes a list of parsed game object and integer values for number of red, green and blue dices and returns the sum of ids of valid games
func getSumOfValidGames(games []Game, red int, green int, blue int) int {
	idSum := 0

	for _, g := range games {
		if g.IsValid(red, green, blue) {
			idSum += g.id
		}
	}
	return idSum
}

func getSumOfPowers(games []Game) int {
	powerSum := 0
	
	for _, game := range games {
		powerSum += game.GetGamePower()	
	}

	return powerSum
}
