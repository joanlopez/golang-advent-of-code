package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	target    = 2020
	inputFile = "input-1"
)

func main() {
	numbers, err := readNumbersFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

loop:
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				fmt.Println(numbers[i] * numbers[j])
				break loop
			}
		}
	}
}

func readNumbersFromFile(filename string) (numbers []int64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}

	return numbers, scanner.Err()
}
