package example

import (
	"sub"
	"pbm"
)


type ExampleSolver struct {

}

func (ExampleSolver) Solve(p pbm.Pbm) sub.Sub {
	s := sub.CreateSub(p.Fleet)
	s.Sub[0] = []int {0}
	s.Sub[1] = []int {2, 1}
	return s
}