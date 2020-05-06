package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := reader()
	numList := getNumList(str)
	printList(insertionSort(numList))
}

func printList(numList []int) {
	for _, val := range numList {
		fmt.Print(val)
	}
}

func insertionSort(numList []int) []int {

	for i := 1; i < len(numList); i++ {
		key := numList[i]
		j := i - 1

		for j >= 0 && numList[j] > key {
			numList[j+1] = numList[j]
			j = j - 1
		}
		numList[j+1] = key
	}
	return numList
}

func getNumList(strNumList string) []int {
	t := strings.Split(strNumList, " ")
	var t2 = []int{}

	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}

func reader() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)
	return str
}
