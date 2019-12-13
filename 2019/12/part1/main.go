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

	steps = 1000
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Vector struct {
	X int
	Y int
	Z int
}

func (v1 Vector) Sum(v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

type Moon struct {
	Pos Vector
	Vel Vector
}

func NewMoon(x, y, z int) *Moon {
	return &Moon{
		Pos: Vector{X: x, Y: y, Z: z},
		Vel: Vector{},
	}
}

func UpdatePosAndVel(moons []*Moon) {
	for _, self := range moons {
		for _, other := range moons {
			UpdateVel(self, other)
		}
	}

	for _, self := range moons {
		UpdatePos(self)
	}
}

func UpdateVel(self, other *Moon) {
	comparison := ComparePos(self, other)
	self.Vel = self.Vel.Sum(comparison)
}

func ComparePos(self, other *Moon) Vector {
	x, y, z := 0, 0, 0
	if self.Pos.X > other.Pos.X {
		x = -1
	} else if self.Pos.X < other.Pos.X {
		x = 1
	}

	if self.Pos.Y > other.Pos.Y {
		y = -1
	} else if self.Pos.Y < other.Pos.Y {
		y = 1
	}

	if self.Pos.Z > other.Pos.Z {
		z = -1
	} else if self.Pos.Z < other.Pos.Z {
		z = 1
	}

	return Vector{X: x, Y: y, Z: z}
}

func UpdatePos(self *Moon) {
	self.Pos = self.Pos.Sum(self.Vel)
}

func TotalEnergy(moons []*Moon) (total int) {
	for _, moon := range moons {
		pot := Abs(moon.Pos.X) + Abs(moon.Pos.Y) + Abs(moon.Pos.Z)
		kin := Abs(moon.Vel.X) + Abs(moon.Vel.Y) + Abs(moon.Vel.Z)
		total += pot * kin
	}
	return
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

	// do N steps
	for i := 0; i < steps; i++ {
		UpdatePosAndVel(moons)
	}

	fmt.Printf("the total energy is %d\n", TotalEnergy(moons))
}
