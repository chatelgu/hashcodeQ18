package main

import (
	"os"
	"logger"
	"pbm"
	"fmt"
	"sub"
)

const TAG = "Main"


func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	args := os.Args[1:]
	name := "dummy" //
	if (len(args) >= 1) {
		name = args[0]
	}

	p := pbm.Parse(name)

	logger.D(TAG, "%s", p)

	sub := sub.Sub{}
	sub.ToFile(name)
}

