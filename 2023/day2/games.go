package day2

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Turn struct {
	r int
	g int
	b int
}

func (t *Turn) New(r int, g int, b int) {
	t.r = r
	t.g = g
	t.b = b
}

type Game struct {
	id    int
	turns []*Turn
}

func (g *Game) New(id int, turns []*Turn) {
	g.id = id
	g.turns = turns
}

func (g *Game) ToString() {
	fmt.Println("id", g.id)
	for i, turns := range g.turns {
		fmt.Println("Turn ", i)
		fmt.Printf("Red: %d, Blue: %d, Green: %d\n", turns.r, turns.g, turns.b)
	}
}

// IsValid Checks is a given game is valid based on the provided number of red, green and blue dices
func (g *Game) IsValid(red int, green int, blue int) bool {
	for _, gameTurn := range g.turns {
		if gameTurn.r > red || gameTurn.g > green || gameTurn.b > blue {
			return false
		}
	}
	return true
}

func (g *Game) ParseTurns(turnStrings []string) []*Turn {
	var turns []*Turn

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

		newTurn := new(Turn)
		newTurn.New(red, blue, green)
		turns = append(turns, newTurn)
	}
	return turns
}

//GetGamePower Returns the product of the power of each dice in the game. A power of a dice is the minimum of dice of the color required for that game to be successful
func (g *Game) GetGamePower() int {
	redPower, greenPower, bluePower := 0,0,0

	for _, t := range g.turns {
		redPower = int(math.Max(float64(redPower), float64(t.r)))
		greenPower = int(math.Max(float64(greenPower), float64(t.g)))
		bluePower = int(math.Max(float64(bluePower), float64(t.b)))
	}
	return redPower * greenPower * bluePower
}
