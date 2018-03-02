package dummy

import (
	"pbm"
	"sub"
	"logger"
	"sort"
)

const TAG = "Dummy"

type vehicule struct {
	rides []int
	step int
	c int
	r int
}

type run []vehicule

var simulation run
var problem pbm.Pbm
var sortedRides []pbm.Ride

func createSimultation(size int) {
	simulation = make([]vehicule, size)
	for i:=0; i<size; i++ {
		simulation[i] = vehicule{make([]int, 0), 0, 0 ,0}
	}
}

func createSortedRides() {
	sortedRides = make([]pbm.Ride, len(problem.Rides))
	for i,ride := range problem.Rides {
		sortedRides[i] = ride
	}
	sort.Sort(pbm.ByDate(sortedRides))
}

func max(i,j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func (v *vehicule) canTakeRide(ride pbm.Ride) (bool, bool, int) {
	distToGo := pbm.Dist(v.r, v.c, ride.RStart, ride.CStart)
	dist := ride.Dist
	bonus := v.step + distToGo <= ride.TimeStart
	inTime := v.step + distToGo + dist <= ride.TimeFinish
	newstep := max(v.step + distToGo, ride.TimeStart) + dist
	return inTime, bonus && inTime, newstep
}

func (v *vehicule) takeRide(ride pbm.Ride) {
	ok, _, newstep := v.canTakeRide(ride)
	if (ok) {
		v.r = ride.RFinish
		v.c = ride.CFinish
		v.step = newstep
		v.rides = append(v.rides, ride.Index)
	} else {
		logger.E(TAG, "Can't take this ride")
	}
}


func findVehicule(ride pbm.Ride) (bool, vehicule, int) {
	best_vehicule_index := -1
	for i,v := range simulation {
		//logger.D(TAG, "vehicule %d", i)
		ok, bonus, _ := v.canTakeRide(ride)
		//logger.D(TAG, "can take %t %t", ok, bonus)
		if (bonus) {
			return true, v , i
		} else if ok {
			best_vehicule_index = i
		}
	}
	if (best_vehicule_index == -1) {
		return false, vehicule{}, -1
	} else {
		return true, simulation[best_vehicule_index], best_vehicule_index
	}
}


type Dummy struct {

}

func (Dummy) Solve(p pbm.Pbm) sub.Sub {
	problem = p
	s := sub.CreateSub(p.Fleet)
	createSimultation(p.Fleet)
	createSortedRides()

	for i,ride := range sortedRides {
		logger.V(TAG, "ride %d : %s", i, ride)
		ok, v, i := findVehicule(ride)
		if ok {
			logger.V(TAG, "ride taken by v")
			v.takeRide(ride)
			logger.V(TAG, "rides", v.rides)
			simulation[i] = v
		}
	}

	for i,v := range simulation {
		s.Sub[i] = v.rides
		logger.V(TAG, "rides ", v.rides)
	}

	return s
}
