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
	items, err := readNumbersFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var validCount int64
	for _, item := range items {
		if item.isValid() {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type item struct {
	minOccurrences int64
	maxOccurrences int64

	letter   rune
	password string
}

func (i item) isValid() bool {
	var occurrences int64

	for _, c := range i.password {
		if i.letter == c {
			occurrences++
		}
	}

	return occurrences >= i.minOccurrences && occurrences <= i.maxOccurrences
}

func readNumbersFromFile(filename string) (items []item, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chunks := strings.Split(scanner.Text(), " ")

		minMaxChunks := strings.Split(chunks[0], "-")
		minOccurrences, err := strconv.ParseInt(minMaxChunks[0], 10, 64)
		if err != nil {
			return nil, err
		}
		maxOccurrences, err := strconv.ParseInt(minMaxChunks[1], 10, 64)
		if err != nil {
			return nil, err
		}

		items = append(items, item{
			minOccurrences: minOccurrences,
			maxOccurrences: maxOccurrences,
			letter:         rune(chunks[1][0]),
			password:       chunks[2],
		})
	}

	return
}
