package main

import (
	"logger"
	"pbm"
	"fmt"
	"algo/dummy"
	"sub"
	"algo/example"
)

const TAG = "Main"


type Solver interface {
	Solve(pbm pbm.Pbm) sub.Sub
}

type algos struct {
	name string
	s Solver
}


func main() {
	fmt.Printf("Hello\n")
	logger.Level = logger.DEBUG

	names := []algos {
		algos{"a_example", example.ExampleSolver{}},
		algos{"b_should_be_easy", dummy.Dummy{}},
		algos{"c_no_hurry",dummy.Dummy{}},
		algos{"d_metropolis",dummy.Dummy{}},
		algos{"e_high_bonus",dummy.Dummy{}},
		algos{"b_should_be_easy", dummy.DummyTmp{}},
		algos{"c_no_hurry",dummy.DummyTmp{}},
		algos{"d_metropolis",dummy.DummyTmp{}},
		algos{"e_high_bonus",dummy.DummyTmp{}},
	}

	for _, algo := range names {
		p := pbm.Parse(algo.name)

		logger.D(TAG, "%s", p)

		sub := algo.s.Solve(p)

		sub.ToFile(algo.name, p)
	}


}

