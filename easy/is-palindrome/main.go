package main

import "strconv"

//Check First and Last loop to center
func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func isPalindromeint(x int) bool {

	str := strconv.Itoa(x)
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
