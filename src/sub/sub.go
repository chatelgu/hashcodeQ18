package sub

import (
	"fileutil"
	"fmt"
)

const TAG = "SUB"


type Sub struct {
	Sub [][]int // list of list of rides id
}

func (s *Sub) ToFile(name string) {
	score := 0
	writer := fileutil.CreateSubFile(name, score)

	//fmt.Fprint(writer, "0")
	for i:=0; i<len(s.Sub); i++ {
		vehicleRides := s.Sub[i]
		M := len(vehicleRides)
		fmt.Fprint(writer, M, " ")
		for j:=0; j<M; j++ {
			fmt.Fprint(writer, vehicleRides[j], " ")
		}
		fmt.Fprintln(writer, "")
	}

	writer.Flush()
}