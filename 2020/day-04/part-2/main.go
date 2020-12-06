package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input-2"

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

	var passports []*Passport
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
	isInvalid  bool
	fieldCount int
}

func NewPassport() *Passport {
	return &Passport{}
}

func (p *Passport) SetInfo(rawData string) {
	chunks := strings.Split(rawData, fieldSeparator)
	for _, chunk := range chunks {
		keyVal := strings.Split(chunk, keyValSeparator)

		if IsAcceptedField(keyVal[0]) {
			p.SetField(keyVal[0], keyVal[1])
		}
	}
}

func (p *Passport) SetField(key, val string) {
	if p.isInvalid {
		return
	}

	switch key {
	case "byr":
		p.SetByr(val)
	case "iyr":
		p.SetIyr(val)
	case "eyr":
		p.SetEyr(val)
	case "hgt":
		p.SetHgt(val)
	case "hcl":
		p.SetHcl(val)
	case "ecl":
		p.SetEcl(val)
	case "pid":
		p.SetPid(val)
	}
}

func (p *Passport) SetByr(val string) {
	n, err := strconv.Atoi(val)
	p.isInvalid = (err != nil) || (n < 1920) || (n > 2002)
	p.fieldCount++
}

func (p *Passport) SetIyr(val string) {
	n, err := strconv.Atoi(val)
	p.isInvalid = (err != nil) || (n < 2010) || (n > 2020)
	p.fieldCount++
}

func (p *Passport) SetEyr(val string) {
	n, err := strconv.Atoi(val)
	p.isInvalid = (err != nil) || (n < 2020) || (n > 2030)
	p.fieldCount++
}

func (p *Passport) SetHgt(val string) {
	p.isInvalid = true
	if len(val) == 4 {
		n, err := strconv.Atoi(val[:2])
		if n >= 59 && n <= 76 && err == nil && val[2:] == "in" {
			p.isInvalid = false
		}
	}

	if len(val) == 5 {
		n, err := strconv.Atoi(val[:3])
		if n >= 150 && n <= 193 && err == nil && val[3:] == "cm" {
			p.isInvalid = false
		}
	}

	p.fieldCount++
}

func (p *Passport) SetHcl(val string) {
	match, _ := regexp.MatchString("^#[a-f0-9]{6}$", val)
	p.isInvalid = !match
	p.fieldCount++
}

func (p *Passport) SetEcl(val string) {
	if !(val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth") {
		p.isInvalid = true
	}
	p.fieldCount++
}

func (p *Passport) SetPid(val string) {
	_, err := strconv.Atoi(val)
	if len(val) != 9 || err != nil {
		p.isInvalid = true
	}
	p.fieldCount++
}

func (p *Passport) IsValid() bool {
	return p.fieldCount == len(fields) && !p.isInvalid
}
