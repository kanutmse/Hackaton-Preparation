package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// type Sortint []int

//Default Sort
// func (s Sortint) Len() int {
// 	return len(s)
// }

// func (s Sortint) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s Sortint) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

func miniMaxSum(arr []int) {
	// Write your code here

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Printf("%d %d", maxSortedSum(arr), minSortedSum(arr))

}

func minSortedSum(arr []int) int {

	// fmt.Print(arr[1:])
	var sum int
	for _, v := range arr[1:] {
		sum += v
	}

	return sum
}

func maxSortedSum(arr []int) int {
	// fmt.Print(arr[:len(arr)-1])
	var sum int
	for _, v := range arr[:len(arr)-1] {
		sum += v
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
