package main

import (
	"os"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/app"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide directory in which csv files present")
	}

	directory := os.Args[1]

	app.RunFromFiles(directory)
}
