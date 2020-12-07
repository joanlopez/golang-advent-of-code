package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input-1"

	goal = "shiny gold"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	bagConversionMap := buildConversionMapFromFile(file)

	queue := bagConversionMap[goal]
	colors := make(map[string]struct{})

	for queue.Len() > 0 {
		find := queue.Remove(queue.Front()).(Conversion)
		colors[find.from] = struct{}{}

		if _, ok := bagConversionMap[find.from]; ok {
			queue.PushBackList(bagConversionMap[find.from])
		}
	}

	fmt.Println(len(colors))

}

type Conversion struct {
	from string
	cost int
}

func buildConversionMapFromFile(file io.Reader) map[string]*list.List {
	bagConversionMap := make(map[string]*list.List)
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

		if _, ok := bagConversionMap[to]; !ok {
			bagConversionMap[to] = list.New()
		}

		bagConversionMap[to].PushBack(Conversion{from: from, cost: cost})

		for _, relation := range relations[1:] {
			chunks = strings.Split(relation, " ")

			to = fmt.Sprintf("%s %s", chunks[2], chunks[3])
			cost, err = strconv.Atoi(chunks[1])
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := bagConversionMap[to]; !ok {
				bagConversionMap[to] = list.New()
			}

			bagConversionMap[to].PushBack(Conversion{from: from, cost: cost})
		}
	}

	return bagConversionMap
}
