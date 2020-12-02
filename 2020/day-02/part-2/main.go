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
	inputFile = "input-2"
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
	firstPos  int64
	secondPos int64

	letter   rune
	password string
}

func (i item) isValid() bool {
	if i.positionOverflow() {
		return false
	}

	return (i.isCharAtPos(int(i.firstPos-1)) && !i.isCharAtPos(int(i.secondPos-1))) || (!i.isCharAtPos(int(i.firstPos-1)) && i.isCharAtPos(int(i.secondPos-1)))
}

func (i item) isCharAtPos(pos int) bool {
	return rune(i.password[pos]) == i.letter
}

func (i item) positionOverflow() bool {
	return int(i.secondPos-1) >= len(i.password)
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
		firstPos, err := strconv.ParseInt(minMaxChunks[0], 10, 64)
		if err != nil {
			return nil, err
		}
		secondPos, err := strconv.ParseInt(minMaxChunks[1], 10, 64)
		if err != nil {
			return nil, err
		}

		items = append(items, item{
			firstPos:  firstPos,
			secondPos: secondPos,
			letter:    rune(chunks[1][0]),
			password:  chunks[2],
		})
	}

	return
}
