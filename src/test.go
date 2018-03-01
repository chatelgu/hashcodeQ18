package main

import (
	"os"
	"logger"
	"pbm"
	"fmt"
)

const TAG = "Test"


func main() {
	fmt.Printf("Test\n")
	logger.Level = logger.DEBUG

	args := os.Args[1:]
	name := "a_example" //
	if (len(args) >= 1) {
		name = args[0]
	}

	p := pbm.Parse(name)

	logger.D(TAG, "%s", p)

	vehicle := pbm.Vehicle{0,0, 0, false}
	ride := pbm.BuildRide(2,2,4,4)
	dist := vehicle.DistToRideStart(ride)

	logger.D(TAG, "dist=%d", dist)
}

