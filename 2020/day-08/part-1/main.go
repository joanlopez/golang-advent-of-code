package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input-1"
)

func main() {
	lines, err := readLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ip, acc := 0, 0
	visitedInstructions := make(map[int]struct{})

	for {
		if _, visited := visitedInstructions[ip]; visited {
			break
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

	fmt.Println(acc)
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
