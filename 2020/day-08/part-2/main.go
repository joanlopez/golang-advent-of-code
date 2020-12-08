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
	inputFile = "input-2"
)

func main() {
	lines, err := readLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	for l := 0; l < len(lines); l++ {
		line := lines[l]

		op := strings.Split(line, " ")[0]
		if op == "nop" {
			lines[l] = strings.Replace(lines[l], "nop", "jmp", 1)
		} else if op == "jmp" {
			lines[l] = strings.Replace(lines[l], "jmp", "nop", 1)
		} else {
			continue
		}

		acc, completed := execute(lines)
		if completed {
			fmt.Println(acc)
			break
		}

		lines[l] = line
	}

}

func execute(lines []string) (int, bool) {
	ip, acc := 0, 0
	visitedInstructions := make(map[int]struct{})

	for {
		if ip == len(lines)-1 {
			return acc, true
		}

		if _, visited := visitedInstructions[ip]; visited {
			return acc, false
		}

		visitedInstructions[ip] = struct{}{}

		chunks := strings.Split(lines[ip], " ")
		op := chunks[0]
		arg, err := strconv.ParseInt(chunks[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		switch op {
		case "nop":
			ip++
		case "jmp":
			ip += int(arg)
		case "acc":
			acc += int(arg)
			ip++
		}

	}
}

func readLinesFromFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
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
