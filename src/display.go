package main

import (
	"fmt"
	"logger"
	"os"
	"pbm"
)

const TAG2 = "MAIN"

func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	args := os.Args[1:]
	name := "dummy" //
	if (len(args) >= 1) {
		name = args[0]
	}

	p := pbm.Parse(name)

	logger.D(TAG2, "%s", p)
	for _, r := range p.Rides {
		logger.D(TAG2, "%s", r)
	}
}