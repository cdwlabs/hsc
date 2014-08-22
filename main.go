package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	log.SetOutput(os.Stderr)

	return 0
}
