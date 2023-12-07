package day1

// https://adventofcode.com/2023/day/1

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// A number map to replace number strings with strings having the number and the first and last character of the spelling of the number.
// This helps avoid edge cases like "eighthree"
var numberMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func parseInput(filename string) []string {
	var inputLines []string
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}
	return inputLines
}

func getDigitFromString(input string) int {
	i := 0
	j := len(input) - 1
	ones := math.MaxInt32
	tens := math.MaxInt32

	for i < len(input) && j >= 0 {
		if tens == math.MaxInt32 && input[i] >= 48 && input[i] <= 57 {
			tens, _ = strconv.Atoi(string(input[i]))
		}
		i++

		if ones == math.MaxInt32 && input[j] >= 48 && input[j] <= 57 {
			ones, _ = strconv.Atoi(string(input[j]))
		}
		j--
	}
	return (tens * 10) + ones
}

func getRealNumberSum(inputLines []string) int {
	sum := 0

	for _, line := range inputLines {
		sum = sum + getDigitFromString(line)
	}
	return sum
}

func getParsedNumberSum(inputLines []string) int {
	sum := 0
	
	for _, line := range inputLines {
		for word := range numberMap {
			if idx := strings.Index(line, word); idx != -1 {
				line = strings.Replace(line, word, string(numberMap[word]), -1)
			}
		}
		sum += getDigitFromString(line)
	}
	return sum
}
