package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"

	numberOfPhases = 100
	numberOfDigits = 8

	inputTimes  = 10000
	inputOffset = 5975483
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

	inputSignal := intArrayTimes(charSliceToInt(chars), inputTimes)[inputOffset:]

	for x := 0; x < numberOfPhases; x++ {
		outputSignal := make([]int, 0, len(inputSignal))
		sum := sumIntSlice(inputSignal)
		for i := range inputSignal {
			outputSignal = append(outputSignal, ((sum%10)+10)%10)
			sum -= inputSignal[i]

		}
		inputSignal = outputSignal
	}
	fmt.Println(inputSignal[:numberOfDigits])
}

func sumIntSlice(nums []int) (sum int) {
	for n := range nums {
		sum += nums[n]
	}
	return
}

func intArrayTimes(nums []int, times int) []int {
	newArray := make([]int, 0, len(nums)*times)
	for i := 0; i < times; i++ {
		for n := range nums {
			newArray = append(newArray, nums[n])
		}
	}
	return newArray
}

func charSliceToInt(chars []string) []int {
	nums := make([]int, 0, len(chars))
	for c := range chars {
		n, _ := strconv.Atoi(chars[c])
		nums = append(nums, n)
	}
	return nums
}
