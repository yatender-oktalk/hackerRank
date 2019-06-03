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

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	result := strings.Split(str, " ")

	var sum int
	var maxSum int
	var tempSum int
	for i := 0; i < totalNum; i++ {
		num, err2 := strconv.Atoi(result[i])
		checkError(err2)

		// tempSum += num
		// fmt.Println("tempSum", tempSum)

		sum = getMaximum(getMaximum(tempSum+num, num), 0)
		// if tempSum < num || tempSum < 0 {
		// 	sum = 0
		// 	// fmt.Println("index", i)
		// 	// fmt.Println("num", num)
		// 	// fmt.Println("sum", sum)

		// } else {
		// 	sum += num

		// 	// fmt.Println("index", i)
		// 	// fmt.Println("sum", sum)
		// 	// fmt.Println("num", num)
		// }

		tempSum = sum

		maxSum = getMaximum(maxSum, sum)

		fmt.Println("max sum", maxSum)
	}

	fmt.Println(maxSum)

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

func getMaximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
