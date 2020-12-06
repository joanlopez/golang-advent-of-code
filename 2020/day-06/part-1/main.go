package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input-1"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	totalSum := 0
	currGroup := make(map[rune]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			totalSum += len(currGroup)
			currGroup = make(map[rune]struct{})
			continue
		}

		for _, char := range line {
			currGroup[char] = struct{}{}
		}
	}

	// Sum the last one
	totalSum += len(currGroup)

	fmt.Println(totalSum)
}
