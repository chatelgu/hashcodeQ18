package pbm

import (
	"fileutil"
	"fmt"
	"bufio"
)

const TAG = "PBM"



type Ride struct {
	RStart int
	CStart int
	RFinish int
	CFinish int
	TimeStart int
	TimeFinish int
	Dist int
	Possible bool
	StartOnTime bool
}

func abs (i int) int {
	if i<0 {
		return -i
	}
	return i
}

func Dist(c1, r1, c2, r2 int ) int {
	return abs(r1 - r2) + abs(c1 -c2)
}
func (r Ride) cmpDist() int {
	return abs(r.RStart - r.RFinish) + abs(r.CStart -r.CFinish)
}

func (r Ride) cmpPossible() bool {
	return r.Dist <= (r.TimeFinish - r.TimeStart)
}


type Vehicle struct {
	I int // index
	R int
	C int
	OnRide bool
}

func (v Vehicle) distToRideStart(r Ride) int {
	return abs(v.R - r.RStart) + abs(v.C -r.CStart)
}

func (v Vehicle) getClosestFreeRide(rides []Ride) Ride {
	closestDist := v.distToRideStart(rides[0])
	closest := rides[0]
	for i := range rides {
		ride := rides[i]
		dist := v.distToRideStart(ride)
		if dist < closestDist {
			closest = ride
			closestDist = dist
		}
	}
	return closest
}

type Pbm struct {
	Name string
	Row int
	Column int
	Fleet int
	Bonus int
	Step int
	Rides []Ride
}

// constructors

//func BuildRide()

func Parse(name string) Pbm {
	var p Pbm
	p.Name = name

	// open the file
	reader := fileutil.OpenPbmFile(name)

	// read data and construct pbm
	var ride int
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"%d %d %d %d %d %d", &p.Row, &p.Column , &p.Fleet, &ride, &p.Bonus, &p.Step)
	p.Rides = make([]Ride, ride)
	for i := 0; i<ride; i++ {
		var r Ride
		scanner.Scan()
		fmt.Sscanf(scanner.Text(),"%d %d %d %d %d %d", &r.RStart, &r.CStart, &r.RFinish, &r.CFinish, &r.TimeStart, &r.TimeFinish)
		r.Dist = r.cmpDist()
		r.Possible = r.cmpPossible()
		p.Rides[i] = r
	}

	return p
}

// pretty print

func (p Pbm) String() string {
	return fmt.Sprintf("Pbm: %s map[%d,%d] vehicules %d bonus %d step %d rides %d", p.Name, p.Row, p.Column, p.Fleet, p.Bonus, p.Step, len(p.Rides))
}

func (r Ride) String() string {
	return fmt.Sprintf("Ride: [%d,%d] to [%d, %d] Start %d Finish %d Dist %d Possible %t", r.RStart, r.CStart, r.RFinish, r.CFinish, r.TimeStart, r.TimeFinish, r.Dist, r.Possible)
}
