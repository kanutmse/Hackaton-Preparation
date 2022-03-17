package main

import (
	"fmt"
	"math"
)

// func reverseString(s string) string {

// 	reverseString := strings.Split(s, "")

// 	sort.Sort(sort.Reverse(sort.StringSlice(reverseString)))

// 	return strings.Join(reverseString, "")
// }

//More Soulution from other ppl
//https://leetcode.com/problems/longest-palindromic-substring/discuss/?currentPage=1&orderBy=hot&query=&tag=go

func getPalindrome(s1 string) string {

	var startStr string
	var middle string
	var endStr string

	if len(s1)%2 != 0 {
		middle = string(s1[len(s1)/2])
	}
	for i, j := 0, len(s1)-1; i < len(s1)/2; i, j = i+1, j-1 {

		if s1[i] != s1[j] {
			return ""
		}
		startStr = startStr + string(s1[i])
		endStr = string(s1[i]) + endStr
	}

	return startStr + middle + endStr
}

func expandAroundCenter(s string, l, r int) int {

	//Check Possible Palindrome
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func longestPalindrome(s string) string {

	//why though why not ?
	if s == "" {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {

		len1 := expandAroundCenter(s, i, i)

		len2 := expandAroundCenter(s, i, i+1)

		curLen := int(math.Max(float64(len1), float64(len2)))

		if curLen > end-start {
			start = i - (curLen-1)/2
			end = i + curLen/2
		}
	}

	return s[start : end+1]
}

func longestPalindromeBruteForce(s string) string {

	//Skip non processing len
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	//init value
	lStart, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}
		// skip duplicates
		k, j := i, i
		for k < len(s)-1 && s[k] == s[k+1] {
			k++
		}
		// expand around centre
		i = k + 1
		for k < len(s)-1 && j > 0 && s[k+1] == s[j-1] {
			k++
			j--
		}
		newLen := k - j + 1
		if newLen > maxLen {
			maxLen = newLen
			lStart = j
		}
	}
	return s[lStart : lStart+maxLen]
}

func main() {

	str := "12321"
	length := int(math.Ceil(float64(len(str)) / 2.0))
	lstr := str[:length]
	rstr := str[length-1:]
	fmt.Printf("%s %s\n", lstr, rstr)
	str2 := "1221"

	// length = int(math.Ceil(float64(len(str)) / 2.0))
	// lstr = str[:length]
	// rstr = str[length:]
	// fmt.Printf("%s %s\n", lstr, rstr)
	str3 := "1221"

	fmt.Printf("%s\n", longestPalindrome("babad"))
	fmt.Printf("%s\n", longestPalindrome(str2))
	fmt.Printf("%s\n", longestPalindrome(str3))
}
