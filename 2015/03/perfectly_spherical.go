package perfectly_spherical

import "fmt"

// Very ugly code alert!!!

func HousesWithPresent(route string) int {
	houses := make(map[string]int)
	x := 0
	y := 0

	coords := fmt.Sprintf("%v,%v", x, y)
	houses[coords]++

	for _, c := range route {
		if c == '^' {
			y++
		}

		if c == 'v' {
			y--
		}

		if c == '>' {
			x++
		}

		if c == '<' {
			x--
		}

		coords := fmt.Sprintf("%v,%v", x, y)
		houses[coords]++
	}

	return len(houses)
}

func HousesWithPresentRoboted(route string) int {
	houses := make(map[string]int)
	xSanta := 0
	ySanta := 0

	xRobo := 0
	yRobo := 0

	coords := fmt.Sprintf("%v,%v", xSanta, ySanta)
	houses[coords]++

	coords = fmt.Sprintf("%v,%v", xRobo, yRobo)
	houses[coords]++

	for i, c := range route {
		if c == '^' {
			if i % 2 == 0 {
				ySanta++
			} else {
				yRobo++
			}
		}

		if c == 'v' {
			if i % 2 == 0 {
				ySanta--
			} else {
				yRobo--
			}
		}

		if c == '>' {
			if i % 2 == 0 {
				xSanta++
			} else {
				xRobo++
			}
		}

		if c == '<' {
			if i % 2 == 0 {
				xSanta--
			} else {
				xRobo--
			}
		}

		var coords string
		if i % 2 == 0 {
			coords = fmt.Sprintf("%v,%v", xSanta, ySanta)
		} else {
			coords = fmt.Sprintf("%v,%v", xRobo, yRobo)
		}
		houses[coords]++
	}

	return len(houses)
}
