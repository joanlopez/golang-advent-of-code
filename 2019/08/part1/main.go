package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

const (
	inputFile = "input.txt"

	layerWide = 25
	layerTall = 6
	layerSize = layerWide * layerTall
)

func main() {
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)

	fewestZeroes, fewestZeroesLayer := math.MaxInt64, 0
	for layer := 0; (layer * layerSize) < len(content); layer++ {
		zeroesCount := countLayerDigitAppearances(layer, 0, content)

		if zeroesCount < fewestZeroes {
			fewestZeroes = zeroesCount
			fewestZeroesLayer = layer
		}
	}

	fmt.Printf("the layer with fewest zeroes is the layer %d with %d zeroes\n", fewestZeroesLayer, fewestZeroes)

	ones := countLayerDigitAppearances(fewestZeroesLayer, 1, content)
	twos := countLayerDigitAppearances(fewestZeroesLayer, 2, content)

	fmt.Printf("and it contains %d one digits and %d two digits, what results on: %d\n", ones, twos, ones*twos)
}

func countLayerDigitAppearances(layer, digit int, content string) (appearances int) {
	digitStr := strconv.Itoa(digit)
	for i := layer * layerSize; i < (layer+1)*layerSize; i++ {
		if string(content[i]) == digitStr {
			appearances++
		}
	}
	return
}
