package no_math

import (
	"bufio"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

var funcs = map[string]func(dimSeq string) int {"paper": NeededPaper, "ribbon": NeededRibbon}

func min(nums ...int) (min int) {
	min = int(math.MaxInt32) // Could depend on the architecture
	for _, n := range nums {
		if min > n {
			min = n
		}
	}
	return
}

func mins(nums ...int) (min1, min2 int) {
	min1 = int(math.MaxInt32) // Could depend on the architecture
	min2 = int(math.MaxInt32) // Could depend on the architecture
	for _, n := range nums {
		if n < min2 {
			min2 = n
		}

		if n < min1 {
			min2 = min1
			min1 = n
		}
	}
	return
}

func NeededPaper(dimSeq string) int {
	re := regexp.MustCompile("[0-9]+")
	dim := re.FindAllString(dimSeq, -1)
	length, _ := strconv.Atoi(dim[0])
	width, _ := strconv.Atoi(dim[1])
	height, _ := strconv.Atoi(dim[2])
	minSide := min(length*width, length*height, width*height)
	return 2*length*width + 2*width*height + 2*height*length + minSide
}

func NeededRibbon(dimSeq string) int {
	re := regexp.MustCompile("[0-9]+")
	dim := re.FindAllString(dimSeq, -1)
	length, _ := strconv.Atoi(dim[0])
	width, _ := strconv.Atoi(dim[1])
	height, _ := strconv.Atoi(dim[2])
	bowRibbon := length*width*height
	min1, min2 := mins(length, width, height)
	presentRibbon := min1*2 + min2*2
	return bowRibbon + presentRibbon
}

func NeededResourceFromFile(filepath, resource string) (feets int) {
	f, _ := os.Open(filepath)
	defer f.Close()

	r := bufio.NewReader(f)

	feets = 0

	line, _, err := r.ReadLine()
	for err != io.EOF {
		feets += funcs[resource](string(line))
		line, _, err = r.ReadLine()
	}
	return
}