package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile = "input-1"

	preamble = 25
)

func main() {
	numbersSlice, err := readNumbersFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	for i := preamble; i < len(numbersSlice); i++ {
		numbersMap := make(map[int64]struct{})
		for j := i - preamble; j < i; j++ {
			numbersMap[numbersSlice[j]] = struct{}{}
		}

		var found bool
		for j := i - preamble; j < i; j++ {
			rem := numbersSlice[i] - numbersSlice[j]
			if _, ok := numbersMap[rem]; ok {
				found = true
				break
			}
		}

		if !found {
			fmt.Println(numbersSlice[i])
			break
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
