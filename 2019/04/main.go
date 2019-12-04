package main

import (
	"fmt"
)

const (
	FROM = 130254
	TO   = 678275
)

// It's very ugly, I know it, but it works :D
// It's all about competitive programming 0:)
func main() {
	fmt.Println(countCombinations(FROM, TO))
	fmt.Println(countCombinationsV2(FROM, TO))
}

func countCombinations(from, to int) (count int) {
	for i := from; i <= to; i++ {
		if isPassword(i) {
			count++
		}
	}
	return count
}

func countCombinationsV2(from, to int) (count int) {
	for i := from; i <= to; i++ {
		if isPasswordV2(i) {
			count++
		}
	}
	return count
}

func isPassword(n int) bool {
	n6 := n % 10
	n5 := (n / 10) % 10
	n4 := (n / 100) % 10
	n3 := (n / 1000) % 10
	n2 := (n / 10000) % 10
	n1 := n / 100000
	return (n1 <= n2 && n2 <= n3 && n3 <= n4 && n4 <= n5 && n5 <= n6) &&
		(n1 == n2 || n2 == n3 || n3 == n4 || n4 == n5 || n5 == n6)
}

func isPasswordV2(n int) bool {
	var matching bool
	var matchSize uint8
	var matchSatisfied bool

	n6 := n % 10
	n5 := (n / 10) % 10

	if n6 < n5 {
		return false
	}

	if n6 == n5 {
		matching = true
		matchSize = 2
	}

	n4 := (n / 100) % 10

	if n5 < n4 {
		return false
	}

	if n5 == n4 {
		if matching {
			matchSize++
		} else {
			matching = true
			matchSize = 2
		}
	}

	if n5 > n4 {
		if matching {
			matchSatisfied = true
		}

		matching = false
		matchSize = 0
	}

	n3 := (n / 1000) % 10

	if n4 < n3 {
		return false
	}

	if n4 == n3 {
		if matching {
			matchSize++
		} else {
			matching = true
			matchSize = 2
		}
	}

	if n4 > n3 {
		if matching && matchSize == 2 {
			matchSatisfied = true
		}

		matching = false
		matchSize = 0
	}

	n2 := (n / 10000) % 10

	if n3 < n2 {
		return false
	}

	if n3 == n2 {
		if matching == true {
			matchSize++
		} else {
			matching = true
			matchSize = 2
		}
	}

	if n3 > n2 {
		if matching && matchSize == 2 {
			matchSatisfied =  true
		}

		matching = false
		matchSize = 0
	}

	n1 := n / 100000

	if n2 < n1 {
		return false
	}

	if n2 == n1 {
		if matching == true {
			matchSize++
		} else {
			matching = true
			matchSize = 2
		}
	}

	if n2 > n1 {
		if matching && matchSize == 2 {
			matchSatisfied = true
		}
	}

	return matchSatisfied || matching && matchSize == 2
}
