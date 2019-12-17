package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"

	numberOfPhases = 100
	numberOfDigits = 8
)

func main() {
	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read line by line
	contents, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	chars := strings.Split(string(contents), "")

	inputSignal := charSliceToInt(chars)

	for x := 0; x < numberOfPhases; x++ {
		outputSignal := make([]int, 0, len(inputSignal))
		for i := range inputSignal {
			total := 0
			pattern := buildPattern(i + 1)
			for j := range inputSignal {
				nInput := inputSignal[j]
				nPattern := pattern[(j+1)%len(pattern)]
				total += nInput * nPattern
			}
			outputSignal = append(outputSignal, digit(abs(total), 1))
		}
		inputSignal = outputSignal
	}
	fmt.Println(inputSignal[:numberOfDigits])
}

func charSliceToInt(chars []string) []int {
	nums := make([]int, 0, len(chars))
	for c := range chars {
		n, _ := strconv.Atoi(chars[c])
		nums = append(nums, n)
	}
	return nums
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func buildPattern(repeats int) []int {
	values := []int{0, 1, 0, -1}
	pattern := make([]int, 0, repeats*4)
	for i := range values {
		for j := 0; j < repeats; j++ {
			pattern = append(pattern, values[i])
		}
	}
	return pattern
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
