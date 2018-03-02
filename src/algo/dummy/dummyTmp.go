package dummy

import (
	"pbm"
	"sub"
	"logger"
)

const TAG = "DummyTmp"

func remove(slice []pbm.Ride, s int) []pbm.Ride {
	if s >= len(slice) - 1 {
		return []pbm.Ride{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func addRide(rides []pbm.Ride, v pbm.Vehicle, s sub.Sub) {
	if len(rides)>0 {
		ride := v.GetClosestFreeRide(rides)
		v.Ride = ride
		rides = remove(rides, ride.Id)
		s.Sub = append(s.Sub, []int{ride.Id})
	}
}

func DummyTmp(p pbm.Pbm) sub.Sub {
	logger.D(TAG, "DummyTmp")
	problem = p
	s := sub.CreateSub(p.Fleet)

	rides := p.Rides
	vehicles := make([]pbm.Vehicle, 0)
	for i := 0; i<p.Fleet; i++ {
		v := pbm.BuildVehicle(i, 0, 0)
		vehicles = append(vehicles, v)
	}

	for step := 0; step<p.Step; step++ {
		logger.D(TAG, "Step %d/%d", step, p.Step)
		for i := range vehicles {
			v := vehicles[i]

			if v.Ride == pbm.EmptyRide {
				addRide(rides, v, s)
			} else {
				v.Ride.Step++
				// check if finished
				if v.Ride.Step >= v.Ride.Dist {
					v.Ride = pbm.EmptyRide
					addRide(rides, v, s)
				}
			}
		}
	}

	return s
}