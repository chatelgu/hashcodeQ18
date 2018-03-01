package main

import (
	"fmt"
	"logger"
	"pbm"
	"os"
)

const TAG2 = "MAIN"

func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	//args := os.Args[1:]
	//name := "a_example" //
	//if (len(args) >= 1) {
	//	name = args[0]
	//}

	b := pbm.Parse("b_should_be_easy")
	logger.D(TAG2, "%s", b)

	c := pbm.Parse("c_no_hurry")
	logger.D(TAG2, "%s", c)

	d := pbm.Parse("d_metropolis")
	logger.D(TAG2, "%s", d)

}