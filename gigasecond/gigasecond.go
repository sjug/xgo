package gigasecond

import "time"

const (
	TestVersion = 2
	gigasecond  = 1e9
)

var Birthday = time.Date(2015, 5, 9, 0, 0, 0, 0, time.UTC)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond * time.Second)
}
