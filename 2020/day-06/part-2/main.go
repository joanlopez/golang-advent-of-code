package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input-2"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	totalSum := 0

	currCount, currGroup := 0, make(map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			for _, n := range currGroup {
				if n == currCount {
					totalSum++
				}
			}

			currCount, currGroup = 0, make(map[rune]int)
			continue
		}

		currCount++
		for _, char := range line {
			currGroup[char]++
		}
	}

	// Sum the last one
	totalSum += len(currGroup)

	fmt.Println(totalSum)
}
