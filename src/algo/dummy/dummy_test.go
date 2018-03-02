package dummy

import (
	"testing"
	"pbm"
	"logger"
)

func TestCanTake(t *testing.T) {
	v := vehicule{nil, 0,0, 0}
	rideOk := pbm.BuildRide(100, 100, 200, 200)
	rideOk.TimeStart = 300
	rideOk.TimeFinish = 800
	ok, bonus, newstep := v.canTakeRide(rideOk)
	logger.D(TAG, "%t %t %d", ok, bonus, newstep)
	if !ok {
		t.Fail()
	}
	if !bonus {
		t.Fail()
	}
	if newstep != 500 {
		t.Fail()
	}
}

func TestCanNoBonusTake(t *testing.T) {
	v := vehicule{nil, 0,0, 0}
	rideOk := pbm.BuildRide(100, 100, 200, 200)
	rideOk.TimeStart = 100
	rideOk.TimeFinish = 800
	ok, bonus, newstep := v.canTakeRide(rideOk)
	logger.D(TAG, "%t %t %d", ok, bonus, newstep)
	if !ok {
		t.Fail()
	}
	if bonus {
		t.Fail()
	}
	if newstep != 400 {
		t.Fail()
	}
}


func TestCannotTake(t *testing.T) {
	v := vehicule{nil, 0,0, 0}
	rideOk := pbm.BuildRide(100, 100, 200, 200)
	rideOk.TimeStart = 100
	rideOk.TimeFinish = 300
	ok, bonus, newstep := v.canTakeRide(rideOk)
	logger.D(TAG, "%t %t %d", ok, bonus, newstep)
	if ok {
		t.Fail()
	}
	if bonus {
		t.Fail()
	}
}

