package dummy

import (
	"pbm"
	"sub"
)

func Dummy1perV(p pbm.Pbm) sub.Sub {
	problem = p
	s := sub.CreateSub(p.Fleet)
	//rides := p.Rides
	for i := 0; i<p.Fleet; i++ {
		v := pbm.BuildVehicle(i, 0, 0)
		ride_id := v.I
		s.Sub = append(s.Sub, []int{ride_id})
	}
	return s
}