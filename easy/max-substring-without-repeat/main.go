package main

import (
	"fmt"
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	//Case Length 1
	if len(s) == 1 {
		return 1
	}

	max := 0

	//Case With Replica
	for k, v1 := range s {
		// fmt.Printf("%d %c Start\n", k, v1)
		// // curMAx := 0
		// fmt.Printf("=====================\n")
		// fmt.Printf("%+v loop slice\n", s[k+1:])
		pre := string(v1)
		curMAx := 1
		for _, v2 := range s[k+1:] {

			// fmt.Printf("%d %c Start(Loop 2 )\n", k2, v2)
			// fmt.Printf(" %s  is contain  (%c) \n", pre, v2)
			if strings.Contains(pre, string(v2)) {
				// fmt.Printf("Contain!!!\n")
				if curMAx > max {
					// fmt.Printf("Old Max %d\n", max)
					max = curMAx
					// fmt.Printf("New Max %d\n", max)
					// fmt.Printf("==========END LOOP===========\n")

				}
				break
			}
			curMAx++
			pre = pre + string(v2)
		}

		if curMAx > max {
			max = curMAx
		}

	}

	//All unique case
	// if max == 0 && len(s) > 0 {
	// 	return len(s)
	// }

	return max
}

func lengthOfLongestSubstringOptimized(s string) int {

	hm := make(map[string]int)

	max := 0.00

	var i int
	for j := 0; j < len(s); j++ {

		//Check if this string exist in map
		if _, ok := hm[string(s[j])]; ok {
			//
			i = int(math.Max(float64(i), float64(hm[string(s[j])])))
		}

		max = math.Max(max, float64(j-i+1))
		hm[string(s[j])] = j + 1
	}
	return int(max)

}
func main() {

	fmt.Println(lengthOfLongestSubstringOptimized("abcabca"))
}
