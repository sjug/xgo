package clock

import "fmt"

const TestVersion = 2
const fullDay = 24 * 60

type Clock struct {
	minutes int
}

func Time(hour, minute int) Clock {
	newtotal := hour*60 + minute
	if newtotal < 0 {
		newtotal = newtotal%fullDay + fullDay
	} else {
		newtotal = newtotal % fullDay
	}
	return Clock{
		minutes: newtotal,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c Clock) Add(minutes int) Clock {
	return Time(c.minutes/60, c.minutes%60+minutes)
}
