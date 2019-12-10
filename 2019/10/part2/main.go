package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const (
	inputFile = "input.txt"

	asteroidChar = "#"

	lookForVaporizedAt = 200
)

type Point struct {
	X int
	Y int
}

// Angle returns the angle in degrees adjusted to the problem statement
// in order to start counting degrees from top instead of right, that's
// why it does an internal rotation before calculating the final angle
func Angle(center, to Point) float64 {
	// considering the center as the (0,0)
	p := Point{X: to.X - center.X, Y: to.Y - center.Y}
	// remember a rotate by 90 degrees is (x' = -y,  y' = x)
	degrees := math.Atan2(float64(p.X), float64(-p.Y)) * 180 / math.Pi

	if degrees < 0 {
		return 360 + degrees
	}

	return degrees
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
	var bestSightLines map[float64][]Point
	for _, self := range asteroids {
		sightLines := make(map[float64][]Point)
		for _, other := range asteroids {
			if self == other {
				continue
			}

			angle := Angle(self, other)
			sightLines[angle] = append(sightLines[angle], other)
		}
		reachCount := len(sightLines)
		if reachCount > bestReachable {
			bestAsteroid = self
			bestReachable = reachCount
			bestSightLines = sightLines
		}
	}
	fmt.Printf("The best asteroid (%d,%d) can reach %d asteroids\n", bestAsteroid.X, bestAsteroid.Y, bestReachable)

	sortSightLines(bestAsteroid, bestSightLines)

	angles := make(sort.Float64Slice, 0, len(bestSightLines))
	for angle := range bestSightLines {
		angles = append(angles, angle)
	}

	sort.Sort(angles)

	var idx int
	var vaporizedIdx int
	var vaporizedAt Point
	for vaporizedIdx < lookForVaporizedAt {
		currAngle := angles[idx%len(angles)]
		if len(bestSightLines[currAngle]) == 0 {
			idx++
			continue
		}

		vaporizedAt = bestSightLines[currAngle][0]
		bestSightLines[currAngle] = bestSightLines[currAngle][1:]
		vaporizedIdx++
		idx++
	}

	fmt.Printf("The asteroid (%d,%d) was the %dth vaporized\n", vaporizedAt.X, vaporizedAt.Y, lookForVaporizedAt)
}

func sortSightLines(from Point, sightLines map[float64][]Point) {
	for _, sl := range sightLines {
		sort.SliceStable(sl, func(i, j int) bool { return calculateDist(from, sl[i]) < calculateDist(from, sl[j]) })
	}
}

func calculateDist(from, to Point) float64 {
	xDist := to.X - from.X
	yDist := to.Y - from.Y

	return math.Sqrt(float64(xDist*xDist) + float64(yDist*yDist))
}
