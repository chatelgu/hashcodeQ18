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
	logger.D(TAG2, "%d", a.MaxScore())

	b := pbm.Parse("b_should_be_easy")
	logger.D(TAG2, "%s", b)
	logger.D(TAG2, "%d", b.MaxScore())

	c := pbm.Parse("c_no_hurry")
	logger.D(TAG2, "%s", c)
	logger.D(TAG2, "%d", c.MaxScore())

	d := pbm.Parse("d_metropolis")
	logger.D(TAG2, "%s", d)
	logger.D(TAG2, "%d", d.MaxScore())

	e := pbm.Parse("e_high_bonus")
	logger.D(TAG2, "%s", e)
	logger.D(TAG2, "%d", e.MaxScore())

	logger.D(TAG2, "max score total : %d",a.MaxScore() + b.MaxScore() + c.MaxScore() + d.MaxScore() + e.MaxScore())
}