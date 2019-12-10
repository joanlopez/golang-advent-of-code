package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	inputFile = "input.txt"

	asteroidChar = "#"
)

type Point struct {
	X int
	Y int
}

// Angle returns the angle in degrees
func Angle(from, to Point) float64 {
	return math.Atan2(float64(from.Y-to.Y), float64(from.X-to.X)) * 180 / math.Pi
}

func main() {
	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read line by line
	// store the position of all asteroids
	var y int
	var asteroids []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if string(char) == asteroidChar {
				asteroids = append(asteroids, Point{X: x, Y: y})
			}
		}
		y++
	}

	var bestAsteroid Point
	var bestReachable int
	for _, self := range asteroids {
		sightLines := make(map[float64]int)
		for _, other := range asteroids {
			if self == other {
				continue
			}

			sightLines[Angle(self, other)]++
		}
		reachCount := len(sightLines)
		if reachCount > bestReachable {
			bestAsteroid = self
			bestReachable = reachCount
		}
	}
	fmt.Printf("The best asteroid (%d,%d) can reach %d asteroids\n", bestAsteroid.X, bestAsteroid.Y, bestReachable)
}
