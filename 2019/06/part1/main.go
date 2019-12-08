package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"

	separator = ")"
)

func main() {
	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vertices := make(map[string]bool)
	adjacencyLists := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orbitSplit := strings.Split(scanner.Text(), separator)
		orbited, orbiting := orbitSplit[0], orbitSplit[1]

		adjacencyLists[orbiting] = append(adjacencyLists[orbiting], orbited)
		vertices[orbited] = true
		vertices[orbiting] = true
	}

	var totalConnections int
	for k1 := range vertices {
		for k2 := range vertices {
			if k1 == k2 {
				continue
			}

			if areReachable(k1, k2, adjacencyLists) {
				totalConnections++
			}
		}
	}

	fmt.Println(totalConnections)
}

func areReachable(k1, k2 string, adjacencyLists map[string][]string) bool {
	visited := make(map[string]bool, len(adjacencyLists))

	queue := make([]string, 0)

	visited[k1] = true
	queue = append(queue, k1)

	for len(queue) > 0 {
		k1 = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, adj := range adjacencyLists[k1] {
			if adj == k2 {
				return true
			}

			if !visited[adj] {
				visited[adj] = true
				queue = append(queue, adj)
			}
		}
	}

	return false
}
