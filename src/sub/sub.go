package sub

import (
	"fileutil"
	"fmt"
	"pbm"
	"logger"
	"os"
)

const TAG = "SUB"


type Sub struct {
	Sub [][]int // list of list of rides id
}

func (s *Sub) ToFile(name string, p pbm.Pbm) {
	score := 0
	writer := fileutil.CreateSubFile(name, score)

	//fmt.Fprint(writer, "0")
	for i:=0; i<len(s.Sub); i++ {
		vehicleRides := s.Sub[i]
		M := len(vehicleRides)
		fmt.Fprint(writer, M, " ")
		for j:=0; j<M; j++ {
			rideId := vehicleRides[j]
			ride := p.Rides[rideId]
			score += ride.Dist
			if ride.StartOnTime {
				score += p.Bonus
			}
			fmt.Fprint(writer, rideId, " ")
		}
		fmt.Fprintln(writer, "")
	}
	writer.Flush()

	os.Rename(name+"_0", name+"_"+string(score))
	logger.D(TAG, "File name=%s : score: %v", name, score)
}