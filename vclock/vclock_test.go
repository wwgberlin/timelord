package vclock

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	const bobby = "bobby"
	v1 := NewVclock("data1")
	v1.Incr(bobby)
	v1.Incr(bobby)
	if v1.m[bobby] != 2 {
		t.Error("Did not increment clock for bobby")
	}
}

func TestGetMostRecentSimple(t *testing.T) {
	const bobby = "bobby"
	v1 := NewVclock("data1")
	v1.Incr(bobby)
	v1.Incr(bobby)
	v2 := NewVclock("data2")
	v2.Incr(bobby)
	if v := GetMostRecent([]Vclock{v1, v2}); v.data != "data1" {
		t.Error("GetMostRecent didn't return the correct vector clock")
	}
}

func TestGetMostRecentDifferentLength(t *testing.T) {
	const bobby = "bobby"
	v1 := NewVclock("data1")
	v1.Incr(bobby)
	v1.Incr(bobby)
	v2 := NewVclock("data2")
	v2.Incr(bobby)
	v2.Incr("lucy")
	if v := GetMostRecent([]Vclock{v1, v2}); v.data != "data2" {
		t.Error("GetMostRecent didn't return the correct vector clock")
	}
}
