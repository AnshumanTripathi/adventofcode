package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrebuchet(t *testing.T) {
	inputLines := parseInput("input.txt")
	assert.Equal(t, 54644, getRealNumberSum(inputLines))
}

func TestTrebuchetPart2(t *testing.T) {
	inputLines := parseInput("input.txt")
	assert.Equal(t, 53348, getParsedNumberSum(inputLines))
}
