package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args)

	programmeName := args[0]

	fmt.Printf("The binary name is: %s \n", programmeName)

	otherArgs := args[1:]

	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}

}
