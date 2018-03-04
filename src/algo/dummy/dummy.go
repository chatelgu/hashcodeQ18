package dummy

import (
	"pbm"
	"sub"
	"logger"
	"sort"
)

const TAGD = "Dummy"

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

type ByDate []pbm.Ride

func (s ByDate) Len() int {
	return len(s)
}
func (s ByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDate) Less(i, j int) bool {
	r1 := s[i]
	r2 := s[j]

	return r1.TimeStart < r2.TimeStart
}

type ByDist []pbm.Ride

func (s ByDist) Len() int {
	return len(s)
}
func (s ByDist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDist) Less(i, j int) bool {
	r1 := s[i]
	r2 := s[j]

	return r1.Dist < r2.Dist
}


func createSortedRides() {
	nbVeryLongride := 0
	sortedRides = make([]pbm.Ride, len(problem.Rides))
	for i,ride := range problem.Rides {
		sortedRides[i] = ride
		if (ride.Dist > 10000) {
			nbVeryLongride++
		}
	}
//	sort.Sort(ByDist(sortedRides))
//	l := len(sortedRides)
//	sortedRides = sortedRides[:l-nbVeryLongride]
	sort.Sort(ByDate(sortedRides))
}

func max(i,j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

type ctrResult struct {
	canTakeRide bool
	withBonus bool
	newStep int
	waintingTime int
	timeToGo int
}

func (c1 ctrResult) isBetter(c2 ctrResult) bool {
	if !c2.canTakeRide {
		return true
	} else if !c1.canTakeRide {
		return false
	} else if c1.withBonus && !c2.withBonus {
		return false
	} else if c2.withBonus && !c1.withBonus {
		return true
	} else if c2.withBonus && c1.withBonus {
		return c2.waintingTime <= c1.waintingTime
	} else { // noBonus for the both
		return c2.timeToGo <= c1.timeToGo
	}
}


func (v *vehicule) canTakeRide(ride pbm.Ride) (ctrResult) {
	var ctrr ctrResult
	ctrr.timeToGo = pbm.Dist(v.r, v.c, ride.RStart, ride.CStart)
	dist := ride.Dist
	ctrr.waintingTime = ride.TimeStart - (v.step + ctrr.timeToGo)
	ctrr.canTakeRide = v.step + ctrr.timeToGo + dist <= ride.TimeFinish
	ctrr.withBonus = ctrr.canTakeRide && v.step + ctrr.waintingTime >= 0
	ctrr.newStep = max(v.step + ctrr.timeToGo, ride.TimeStart) + dist
	return ctrr
}

func (v *vehicule) takeRide(ride pbm.Ride) {
	ctrr := v.canTakeRide(ride)
	if (ctrr.canTakeRide) {
		v.r = ride.RFinish
		v.c = ride.CFinish
		v.step = ctrr.newStep
		v.rides = append(v.rides, ride.Index)
	} else {
		logger.E(TAG, "Can't take this ride")
	}
}


func findVehicule(ride pbm.Ride) (bool, vehicule, int) {
	best_vehicule_index := -1
	best_ctrr := ctrResult{false, false, 0, 0, 0}
	for i,v := range simulation {
		//logger.D(TAGD, "vehicule %d", i)
		ctrr := v.canTakeRide(ride)
		//logger.D(TAGD, "can take %t %t", ok, bonus)
		if ctrr.isBetter(best_ctrr) {
			best_vehicule_index = i
			best_ctrr = ctrr
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
		logger.D(TAGD, "ride %d : %s", i, ride)
		ok, v, i := findVehicule(ride)
		if ok {
			logger.V(TAGD, "ride taken by v")
			v.takeRide(ride)
			logger.V(TAGD, "rides", v.rides)
			simulation[i] = v
		}
	}

	for i,v := range simulation {
		s.Sub[i] = v.rides
		logger.V(TAGD, "rides ", v.rides)
	}

	return s
}
