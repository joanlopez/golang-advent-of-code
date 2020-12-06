package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input-1"

	fieldSeparator  = " "
	keyValSeparator = ":"
)

var fields = map[string]struct{}{"byr": {}, "iyr": {}, "eyr": {}, "hgt": {}, "hcl": {}, "ecl": {}, "pid": {}}

func IsAcceptedField(field string) bool {
	_, ok := fields[field]
	return ok
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var passports []Passport
	currPassport := NewPassport()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, currPassport)
			currPassport = NewPassport()
			continue
		}

		currPassport.SetInfo(line)
	}

	// Append the last one
	passports = append(passports, currPassport)

	var validCount int
	for _, passport := range passports {
		if passport.IsValid() {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type Passport struct {
	fields map[string]string
}

func NewPassport() Passport {
	return Passport{
		fields: make(map[string]string),
	}
}

func (p Passport) SetInfo(rawData string) {
	chunks := strings.Split(rawData, fieldSeparator)
	for _, chunk := range chunks {
		keyVal := strings.Split(chunk, keyValSeparator)

		if IsAcceptedField(keyVal[0]) {
			p.fields[keyVal[0]] = keyVal[1]
		}
	}
}

func (p Passport) IsValid() bool {
	return len(p.fields) == len(fields)
}

func (p Passport) Print() {
	fmt.Printf("%v\n", p.fields)
}
