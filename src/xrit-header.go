package main

import (
	"coms"
	"os"
)

func main() {
	args := os.Args[1:]

	coms.Init(args[0])
}
