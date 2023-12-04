package trebuchet

// https://adventofcode.com/2023/day/1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Trebuchet(filename string) {
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
	sum := 0
	for _, line := range inputLines {
		i := 0
		j := len(line) - 1
		ones := math.MaxInt32
		tens := math.MaxInt32
		for i < len(line) && j >= 0 {
			if tens == math.MaxInt32 && line[i] >= 48 && line[i] <= 57 {
				tens, err = strconv.Atoi(string(line[i]))
				if err != nil {
					log.Fatal(err)
				}
			}
			i += 1
			if ones == math.MaxInt32 && line[j] >= 48 && line[j] <= 57 {
				ones, err = strconv.Atoi(string(line[j]))
				if err != nil {
					log.Fatal(err)
				}
			}
			j -= 1
		}
		sum = sum + ((tens * 10) + ones)
	}
	fmt.Println(sum)
}
