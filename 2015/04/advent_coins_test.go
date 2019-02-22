package advent_coins

import (
	"fmt"
	"regexp"
)

//func LowestMatching(input string) (lowest int) {
//	lowest := 1
//	for {
//		txt := fmt.Sprintf("%v%v", input, lowest)
//		hash := md5.New()
//		h := md5.New()
//		io.WriteString(h, "The fog is getting thicker!")
//		fmt.Printf("%x", h.Sum(nil))
//		lowest++
//	}
//
//}


func test() {
	re := regexp.MustCompile("^0{5}")

	fmt.Println(re.Match([]byte("caca")))
}