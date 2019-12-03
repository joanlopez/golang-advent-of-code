package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func calcFuelV1(mass int) int {
	dividedByThree := mass / 3
	roundedDown := int(math.Floor(float64(dividedByThree)))
	return roundedDown - 2
}

func calcFuelV2(mass int) int {
	dividedByThree := mass / 3
	roundedDown := int(math.Floor(float64(dividedByThree)))
	neededFuel := roundedDown - 2

	extraMass := calcFuelV1(neededFuel)
	for extraMass >= 0 {
		neededFuel += extraMass

		extraMass = calcFuelV1(extraMass)
	}
	return neededFuel
}

func main() {
	var totalSumV1, totalSumV2 int

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		totalSumV1 += calcFuelV1(mass)
		totalSumV2 += calcFuelV2(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[Part One] The needed fuel for all the modules is: %d \n", totalSumV1)
	fmt.Printf("[Part Two] The needed fuel for all the modules and its fuel is: %d \n", totalSumV2)
}
