package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse(x int) int {

	minus := false

	if x < 0 {
		minus = true
	}
	str := strconv.Itoa(x)

	if minus {
		//remove before reverse
		str = str[1:]
	}

	str = reverseStr(str)
	fmt.Println(str)
	str = strings.TrimLeft(str, "0")

	if minus {
		str = "-" + str
	}
	i, _ := strconv.Atoi(str)

	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0
	}
	return i
}

func main() {

	// fmt.Println(reverse(-120))s
	fmt.Println(reverse(1534236469))

}
