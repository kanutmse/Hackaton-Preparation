package main

import (
	"math"
	"strings"
)

//https://www.code-recipe.com/post/string_to_integer Explian
func myAtoi(s string) int {

	s = strings.Trim(s, " ")

	n := len(s)
	// If string is empty return 0
	if n == 0 {
		return 0
	}

	// String index from where the processing should start
	var start int

	// Handling numbers with +/- sign
	signMultiplier := 1
	if s[0] == '-' {
		// Handling for negative numbers
		signMultiplier = -1
		start = 1
	} else if s[0] == '+' {
		// Handling for signed positive number
		start = 1
	}

	// ASCII of '0' = 48
	// s[i] - '0' gives the actual number as an integer
	var res int
	for i := start; i < len(s); i++ {
		// Handling for non numeric values
		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMultiplier
		}

		// Calculate the result for current iteration
		res = res*10 + int(s[i]-'0')

		// Check if result exceeds MinInt32 or MaxInt32
		if res*signMultiplier <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultiplier >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * signMultiplier
}
