package sub

import (
	"fileutil"
	"fmt"
	"pbm"
	"logger"
)

const TAG = "SUB"

func CreateSub(size int) Sub {
	var s Sub
	s.Sub = make([][]int, size)
	return s
}

type Sub struct {
	Sub [][]int // list of list of rides id
}

func (s *Sub) Score(p pbm.Pbm) int {
	score := 0
	for v,rides := range s.Sub {
		logger.V(TAG, "vehicule %d", v)
		step := 0
		r,c := 0,0
		for _,ride_i := range rides {
			logger.V(TAG, "moving to ride %d", ride_i)
			ride := p.Rides[ride_i]
			logger.V(TAG, "ride %s", ride)
			step += pbm.Dist(r, c , ride.RStart, ride.CStart)
			logger.V(TAG, "step %d", step)
			if step <= ride.TimeStart { // ontime
				logger.V(TAG, "We get a bonus")
				score += p.Bonus
				logger.V(TAG, "Waiting %d", ride.TimeStart - step)
				step = ride.TimeStart
				logger.V(TAG, "score : %d", score)
			}
			logger.V(TAG, "step %d", step)
			step += ride.Dist
			if step <= ride.TimeFinish {
				score += ride.Dist
			}
			logger.V(TAG, "step %d", step)
			logger.V(TAG, "score %d", score)
		}
		logger.V(TAG, "score %d", score)
	}
	return score
}

func (s *Sub) ToFile(name string, p pbm.Pbm) {
	score := s.Score(p)
	writer := fileutil.CreateSubFile(name, score)

	//fmt.Fprint(writer, "0")
	for i:=0; i<len(s.Sub); i++ {
		vehicleRides := s.Sub[i]
		M := len(vehicleRides)
		fmt.Fprint(writer, M, " ")
		for j:=0; j<M; j++ {
			rideId := vehicleRides[j]
			fmt.Fprint(writer, rideId, " ")
		}
		fmt.Fprintln(writer, "")
	}
	writer.Flush()

	logger.D(TAG, "File name=%s : score: %v", name, score)
}