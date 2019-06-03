package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Complete the maxSubsetSum function below.
// func maxSubsetSum(arr []int32) int32 {

// }

func main() {
	var totalNum = readNum()
	var target = readNum()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	result := strings.Split(str, " ")

	var sum int
	var currentHead int
	var foundSum = -1
	for i := 0; i < totalNum; i++ {
		num, err2 := strconv.Atoi(result[i])
		checkError(err2)
		sum += num

		for sum > target {
			headIndexVal, err := strconv.Atoi(result[currentHead])
			checkError(err)
			sum -= headIndexVal
			currentHead++
		}

		if sum == target {
			foundSum = 1
			fmt.Print("found sum at index ", currentHead, " to ", i)
			break
		}

	}
	if foundSum == -1 {
		fmt.Println("not found ")
	}

}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readNum() int {
	var totalNum int
	fmt.Scan(&totalNum)
	return totalNum
}

func isGreaterThanTarget(target, sum int) bool {
	return sum > target
}
