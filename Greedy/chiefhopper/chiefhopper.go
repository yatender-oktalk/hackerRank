package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var totalCase int
	fmt.Scan(&totalCase)
	if totalCase < 1 || totalCase > 100000 {
		fmt.Println("error")
	}

	nums := readdata()
	if len(nums) != totalCase {
		fmt.Println("error2")
	}
	fmt.Println(doOperation(nums, totalCase))
}

func doOperation(nums []int, totalCase int) int {

	initialEnergy := nums[totalCase-1]
	fmt.Println("initialEnergy")
	for i := len(nums) - 1; i > 1; i-- {
		eneryDiff := nums[i] - nums[i-1]
		initialEnergy += eneryDiff
	}
	return initialEnergy
}

func stringToNumArray(strArr []string) []int {
	inArr := make([]int, len(strArr))
	for i := range strArr {
		inArr[i], _ = strconv.Atoi(strArr[i])
	}
	return inArr
}

func readdata() []int {
	reader := bufio.NewReader(os.Stdin)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	numArr := strings.Split(text2, " ")
	nums := stringToNumArray(numArr)
	return nums
}
