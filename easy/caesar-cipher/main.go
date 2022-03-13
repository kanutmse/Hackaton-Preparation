package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	encryptstring := []rune(s)

	for i, v := range s {

		for j, u := range alphabet {
			index := 0
			if u == unicode.ToLower(v) {
				index = (j + int(k)) % 26

				if unicode.IsUpper(v) {
					encryptstring[i] = unicode.ToUpper(alphabet[index])

				} else {
					encryptstring[i] = rune(alphabet[index])
				}

			}

		}
	}

	// fmt.Printf("%s \n", string(encryptstring))
	return string(encryptstring)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	if n != int32(len(s)) {
		fmt.Println("Hackker Rank suck ass")
	}

	result := caesarCipher(s, k)

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
