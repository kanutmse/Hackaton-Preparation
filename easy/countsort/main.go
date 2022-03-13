package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countingSort' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findMaxVal(arr []int32) int32 {

	var max int32
	max = math.MinInt32

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
func countingSort(arr []int32) []int32 {
	// Write your code here
	// max := findMaxVal(arr)
	ansArr := make([]int32, 100)
	// fmt.Println(max)
	// fmt.Printf("i\tarr[i]\trst\n")
	for i := 0; i < len(arr); i++ {

		// fmt.Printf("%d\t%d\t%v\n", i, arr[i], ansArr)
		ansArr[arr[i]] = ansArr[arr[i]] + 1
	}

	// fmt.Println(len(ansArr))
	return ansArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := countingSort(arr)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("\n")

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
