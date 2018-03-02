package testscore

import (
	"pbm"
	"sub"
)

type TestSolver struct {

}

func (TestSolver) Solve(p pbm.Pbm) sub.Sub {
	s := sub.CreateSub(p.Fleet)
	s.Sub[0] = make([]int, 20)
	for i:= 0; i<20; i++ {
		s.Sub[0][i] = i
	}
	return s
}