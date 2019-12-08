package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"

	separator = ")"

	youNode   = "YOU"
	santaNode = "SAN"
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
		adjacencyLists[orbited] = append(adjacencyLists[orbited], orbiting)

		vertices[orbited] = true
		vertices[orbiting] = true
	}

	min := math.MaxInt64

	for _, youOrbit := range adjacencyLists[youNode] {
		for _, santaOrbit := range adjacencyLists[santaNode] {
			minPath := minPath(youOrbit, santaOrbit, vertices, adjacencyLists)

			if minPath < min {
				min = minPath
			}
		}
	}

	fmt.Println(min)
}

func minPath(k1, k2 string, vertices map[string]bool, adjacencyLists map[string][]string) int {
	nodes := make(map[string]bool, len(vertices))
	distances := make(map[string]int, len(vertices))

	for k := range vertices {
		distances[k] = math.MaxInt64
		nodes[k] = false
	}

	distances[k1] = 0

	for len(nodes) > 0 {
		node := min(nodes, distances)

		if distances[node] == math.MaxInt64 {
			break
		}

		for _, adj := range adjacencyLists[node] {
			newDist := distances[node] + 1

			if newDist < distances[adj] {
				distances[adj] = newDist
			}

		}

		delete(nodes, node)
	}

	return distances[k2]
}

func min(nodes map[string]bool, distances map[string]int) string {
	min := math.MaxInt64
	var minNode string
	for node := range nodes {
		if min > distances[node] {
			minNode = node
			min = distances[node]
		}
	}
	return minNode
}
