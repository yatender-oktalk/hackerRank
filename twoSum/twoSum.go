package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
	fmt.Println(k)
	fmt.Println(arr)

	var count int32
	for i := 0; i < len(arr); i++ {
		firstNum := arr[i]
		for j := 0; j < len(arr); j++ {
			if firstNum-arr[j] == k {
				count++
			}
		}
	}

	return count
}

func pairsMap(k int32, arr []int32) int32 {
	// make map of array values
	var m = make(map[int32]int)

	// initialize a map
	for _, val := range arr {
		m[val] = 1
	}
	var count int32
	for i := 0; i < len(arr); i++ {
		firstNum := arr[i]
		if m[firstNum-k] == 1 {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	fmt.Println("array size and target sum")
	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	fmt.Println("Enter the array values")
	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairsMap(k, arr)
	fmt.Println(result)
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
