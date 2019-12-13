package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input.txt"

	numAxis = 3
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Vector [3]int

func NewVector(x, y, z int) Vector {
	return [3]int{x, y, z}
}

func (v Vector) Sum(v2 Vector) Vector {
	return [3]int{v[0] + v2[0], v[1] + v2[1], v[2] + v2[2]}
}

func (v Vector) IsZero() bool {
	return v[0] == 0 && v[1] == 0 && v[2] == 0
}

type Moon struct {
	Pos Vector
	Vel Vector
}

func NewMoon(x, y, z int) *Moon {
	return &Moon{
		Pos: NewVector(x, y, z),
		Vel: NewVector(0, 0, 0),
	}
}

// find GCD via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find LCM via GCD (at lest two integers are required)
func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// regex to read moon position from a line
	regex, err := regexp.Compile("[x-z]=[-0-9]*")
	if err != nil {
		log.Fatal(err)
	}

	// read file line by line
	var moons []*Moon
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, 3)

		x, _ := strconv.Atoi(matches[0][0][2:])
		y, _ := strconv.Atoi(matches[1][0][2:])
		z, _ := strconv.Atoi(matches[2][0][2:])

		moons = append(moons, NewMoon(x, y, z))
	}

	// storing initial positions in order to detect cycles
	initialPos := make([]Vector, len(moons))
	for i, m := range moons {
		initialPos[i] = m.Pos
	}

	// find the number of required steps
	cycleLength := make([]int, numAxis)
	for i := 0; i < numAxis; i++ {
		steps := 0
		var cycle bool

		for !cycle {
			for _, self := range moons {
				for _, other := range moons {
					if self == other {
						continue
					}

					if self.Pos[i] > other.Pos[i] {
						self.Vel[i] -= 1
					} else if self.Pos[i] < other.Pos[i] {
						self.Vel[i] += 1
					}
				}

			}

			for _, self := range moons {
				self.Pos[i] += self.Vel[i]
			}

			steps++

			cycle = true
			for m := range initialPos {
				if moons[m].Pos[i] != initialPos[m][i] || !moons[m].Vel.IsZero() {
					cycle = false
					break
				}
			}
		}

		cycleLength[i] = steps
	}

	fmt.Printf("the number of required steps is %d\n", LCM(cycleLength...))
}
