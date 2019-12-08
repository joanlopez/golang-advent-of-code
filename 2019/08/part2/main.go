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

	black = 0
	white = 1
	trans = 2

	blackChar = " "
	whiteChar = "X"
)

func main() {
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	resultLayer := make([]int, layerSize)

	// finding the color for each pixel of the result layer
	for i := 0; i < layerSize; i++ {
		resultLayer[i] = math.MaxInt64

		for layer := 0; (layer * layerSize) < len(content); layer++ {
			pixelAtLayerX, _ := strconv.Atoi(string(content[layer*layerSize+i]))

			// the color of each pixel from the result layer corresponds to the
			// first color (no trans) found from the top-to-bottom layers
			if pixelAtLayerX != trans {
				resultLayer[i] = pixelAtLayerX
				break
			}
		}
	}

	// print result layer
	for i := 0; i < layerTall; i++ {
		for j := 0; j < layerWide; j++ {
			if resultLayer[i*layerWide+j] == white {
				fmt.Print(whiteChar)
			} else {
				fmt.Print(blackChar)
			}
		}
		fmt.Println()
	}
}
