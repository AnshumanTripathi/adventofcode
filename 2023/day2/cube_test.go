package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubeConundrumPart1(t *testing.T) {
	parsedGames := parseGames("input.txt")
	assert.Equal(t, 2237, getSumOfValidGames(parsedGames, 12, 14, 13))
}

func TestCubeConundrumPart2(t *testing.T) {
	parsedGames := parseGames("input.txt")
	assert.Equal(t, 66681, getSumOfPowers(parsedGames))
}
