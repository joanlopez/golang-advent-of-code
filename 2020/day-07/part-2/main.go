package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input-2"

	goal = "shiny gold"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	bagConversionMap := buildConversionMapFromFile(file)
	fmt.Println(findCountFor(goal, bagConversionMap))
}

type Conversion struct {
	to   string
	cost int
}

func buildConversionMapFromFile(file io.Reader) map[string][]Conversion {
	bagConversionMap := make(map[string][]Conversion)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		relations := strings.Split(scanner.Text(), ",")

		chunks := strings.Split(relations[0], " ")

		from := fmt.Sprintf("%s %s", chunks[0], chunks[1])
		to := fmt.Sprintf("%s %s", chunks[5], chunks[6])
		cost, err := strconv.Atoi(chunks[4])
		if err != nil {
			continue
		}

		bagConversionMap[from] = append(bagConversionMap[from], Conversion{to: to, cost: cost})

		for _, relation := range relations[1:] {
			chunks = strings.Split(relation, " ")

			to = fmt.Sprintf("%s %s", chunks[2], chunks[3])
			cost, err = strconv.Atoi(chunks[1])
			if err != nil {
				log.Fatal(err)
			}

			bagConversionMap[from] = append(bagConversionMap[from], Conversion{to: to, cost: cost})
		}
	}

	return bagConversionMap
}

func findCountFor(goal string, conversionMap map[string][]Conversion) (count int) {
	if _, ok := conversionMap[goal]; !ok {
		return 0
	}

	for _, c := range conversionMap[goal] {
		count += c.cost + c.cost*findCountFor(c.to, conversionMap)
	}

	return
}
