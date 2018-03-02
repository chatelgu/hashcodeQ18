package main

import (
	"os"
	"logger"
	"pbm"
	"fmt"
	"algo/dummy"
	"sub"
)

const TAG = "Main"

func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	args := os.Args[1:]
	name := "a_example" //
	if (len(args) >= 1) {
		name = args[0]
	}

	p := pbm.Parse(name)

	logger.D(TAG, "%s", p)

	var sub sub.Sub

	logger.D(TAG, "Dummy" )
	sub = dummy.Dummy(p)
	sub.ToFile(name, p)

	logger.D(TAG, "Dummy1perV" )
	sub = dummy.Dummy1perV(p)
	sub.ToFile(name, p)

//	logger.D(TAG, "DummyTmp" )
//	sub = dummy.DummyTmp(p)
//	sub.ToFile(name, p)
}

