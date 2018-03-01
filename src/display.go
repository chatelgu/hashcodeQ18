package main

import (
	"fmt"
	"logger"
	"pbm"
)

const TAG2 = "MAIN"

func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	a := pbm.Parse("a_example")
	logger.D(TAG2, "%s", a)

	b := pbm.Parse("b_should_be_easy")
	logger.D(TAG2, "%s", b)

	c := pbm.Parse("c_no_hurry")
	logger.D(TAG2, "%s", c)

	d := pbm.Parse("d_metropolis")
	logger.D(TAG2, "%s", d)

}