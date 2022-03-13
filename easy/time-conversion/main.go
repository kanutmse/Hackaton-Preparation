package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here

	if strings.Contains(s, "AM") {

		hr, _ := strconv.Atoi(s[:2])

		if hr >= 12 {
			s = strings.Replace(s, s[:2], "00", 1)
			return s[:len(s)-2]
		}

	}

	if strings.Contains(s, "PM") {

		hr, _ := strconv.Atoi(s[:2])

		if hr >= 12 {
			return s[:len(s)-2]
		}
		hr += 12

		if hr >= 24 {
			s = strings.Replace(s, s[:2], "00", 1)
		} else {
			newHr := strconv.Itoa(hr)

			s = strings.Replace(s, s[:2], newHr, 1)
		}

	}

	return s[:len(s)-2]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Printf("%s\n", result)

	// writer.Flush()
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
