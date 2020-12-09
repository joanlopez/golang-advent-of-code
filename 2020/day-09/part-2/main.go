package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	inputFile = "input-2"

	goal = 22406676
)

func main() {
	numbersSlice, err := readNumbersFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	i, j := 0, 1
	currSum := numbersSlice[i] + numbersSlice[j]
	for {
		if currSum == goal {
			break
		} else if currSum > goal {
			currSum -= numbersSlice[i]
			i++
		} else if currSum < goal {
			j++
			currSum += numbersSlice[j]
		}
	}

	min, max := int64(math.MaxInt64), int64(math.MinInt64)
	for k := i; k <= j; k++ {
		if numbersSlice[k] > max {
			max = numbersSlice[k]
		}

		if numbersSlice[k] < min {
			min = numbersSlice[k]
		}
	}

	fmt.Println(min + max)
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
