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

type game struct {
	id    int
	turns []*turn
}


func (g *game) new(id int, turns []*turn) {
	g.id = id
	g.turns = turns
}

func (g *game) toString() {
	fmt.Println("id", g.id)
	for i, turns := range g.turns {
		fmt.Println("Turn ", i)
		fmt.Printf("Red: %d, Blue: %d, Green: %d\n", turns.r, turns.g, turns.b)
	}
}

func (g *game) isValid(red int, blue int, green int) bool {
	for _, gameTurn := range g.turns {
		if gameTurn.r > red || gameTurn.g > green || gameTurn.b > blue {
			return false
		}
	}
	return true
}


func (g *game) parseTurns(turnStrings []string) []*turn {
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

func cubeConundrum(inputFile string, red int, blue int, green int) int {
	var inputLines []string
	idSum :=0

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
		newGame := new(game)
		parsedTurns := newGame.parseTurns(gameTurns)
	 	newGame.new(id, parsedTurns)
		
		if newGame.isValid(red, blue, green) {
			idSum += id
		}
	}
	return idSum
}
