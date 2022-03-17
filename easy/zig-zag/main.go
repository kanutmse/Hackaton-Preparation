package main

import "fmt"

func convert(s string, numRows int) string {

	//Fuck Off
	if numRows == 1 {
		return s
	}

	// map[rows] string
	m := make(map[int]string)

	isZag := false
	zagNum := numRows - 2
	i := 0
	for _, v := range s {

		//Create New Row Every numsRow are reach

		if isZag {
			//Zag Case == numrows - 2
			m[zagNum] = m[zagNum] + string(v)
			// fmt.Println(string(v))
			zagNum--
			if zagNum == 0 {
				isZag = false
				i = 0
			}

		} else {
			//Zig case just put where it should be

			// fmt.Println(string(v))
			m[i%numRows] = m[i%numRows] + string(v)
			i++

			if i%numRows == 0 && numRows-2 > 0 {
				isZag = true
				zagNum = numRows - 2
			}
		}

	}
	ans := ""
	//Get Ans
	for i := 0; i < numRows; i++ {
		ans = ans + m[i]
	}

	return ans
}

func main() {
	fmt.Println(convert("ABCD", 2))
}
