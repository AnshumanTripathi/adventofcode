package day2

// https://adventofcode.com/2023/day/2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type turn struct {
	r int
	g int
	b int
}

func (t *turn) new(r int, g int, b int) {
	t.r = r
	t.g = g
	t.b = b
}

type Game struct {
	id    int
	turns []*turn
}

func (g *Game) new(id int, turns []*turn) {
	g.id = id
	g.turns = turns
}

func (g *Game) toString() {
	fmt.Println("id", g.id)
	for i, turns := range g.turns {
		fmt.Println("Turn ", i)
		fmt.Printf("Red: %d, Blue: %d, Green: %d\n", turns.r, turns.g, turns.b)
	}
}

func (g *Game) isValid(red int, green int, blue int) bool {
	for _, gameTurn := range g.turns {
		if gameTurn.r > red || gameTurn.g > green || gameTurn.b > blue {
			return false
		}
	}
	return true
}

func (g *Game) parseTurns(turnStrings []string) []*turn {
	var turns []*turn

	for _, t := range turnStrings {
		var red, green, blue int
		die := strings.Split(t, ",")

		for _, d := range die {
			d = strings.TrimSpace(d)
			if strings.Contains(d, "red") {
				if r, err := strconv.Atoi(strings.Split(d, " ")[0]); err != nil {
					log.Fatal(err)
				} else {
					red = r
				}
			}

			if strings.Contains(d, "green") {
				if g, err := strconv.Atoi(strings.Split(d, " ")[0]); err != nil {
					log.Fatal(err)
				} else {
					green = g
				}
			}

			if strings.Contains(d, "blue") {
				if b, err := strconv.Atoi(strings.Split(d, " ")[0]); err != nil {
					log.Fatal(err)
				} else {
					blue = b
				}
			}
		}

		newTurn := new(turn)
		newTurn.new(red, blue, green)
		turns = append(turns, newTurn)
	}
	return turns
}

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
		parsedTurns := newGame.parseTurns(gameTurns)
		newGame.new(id, parsedTurns)
		parsedGames = append(parsedGames, *newGame)
	}
	return parsedGames
}

// getSumOfValidGames Takes a list of parsed game object and integer values for number of red, green and blue dices and returns the sum of ids of valid games
func getSumOfValidGames(games []Game, red int, green int, blue int) int {
	idSum := 0

	for _, g := range games {
		if g.isValid(red, green, blue) {
			idSum += g.id
		}
	}
	return idSum
}
