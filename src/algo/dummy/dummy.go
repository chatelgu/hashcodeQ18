package dummy

import (
	"pbm"
	"sub"
	"logger"
)


type Vehicule struct {
	Rides []int
	Step int
}

func (v *Vehicule) getLastPoint() (int,int) {
	l := len(v.Rides)
	if (l == 0) {
		return 0,0
	} else {
		i := v.Rides[l-1]
		ride := problem.Rides[i]
		return ride.RFinish, ride.CFinish
	}
}


func (v *Vehicule) addRide(i int) bool {
	rf, cf := v.getLastPoint();

	ride := problem.Rides[i]
	logger.D("Dummy", "%s", ride)

	distToGo := pbm.Dist(rf, cf, ride.RStart, ride.CStart)

	if v.Step + distToGo + ride.Dist > problem.Step {
		return false
	} else {
		v.Step += distToGo + ride.Dist
		v.Rides = append(v.Rides, i)
		return true
	}
}

type Run struct {
   l []Vehicule
}


func addRides(run *Run, ride, vehicule int) {
	logger.D("Dummy", "try ride %d to %d", ride, vehicule)

	if ride == len(problem.Rides) {
		logger.D("Dummy", "end of it")
		return
	}
	if vehicule == problem.Fleet {
		logger.D("Dummy", "no more car")
		ride++
	} else if run.l[vehicule].addRide(ride) {
		logger.D("Dummy", "ok assigned")
		ride++
	} else {
		logger.D("Dummy", "next car")
		vehicule++
	}
	logger.D("Dummy", "loop")
	addRides(run, ride, vehicule)
}


func createRun(size int) Run {
	var run Run
	run.l = make([]Vehicule, size)
	return run
}

func (run Run) toSub(s *sub.Sub)  {
	for i,r := range run.l {
		s.Sub[i] = r.Rides
	}
}

var problem pbm.Pbm


func Dummy(p pbm.Pbm) sub.Sub {
	problem = p
	s := sub.CreateSub(p.Fleet)
	run := createRun(p.Fleet)
	addRides(&run, 0,0)
	run.toSub(&s)
	return s
}
