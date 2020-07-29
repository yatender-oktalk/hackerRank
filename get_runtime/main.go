package main

import (
	"log"
	"runtime"
)

func main() {
	log.Printf("%s", runtime.Version())
}
