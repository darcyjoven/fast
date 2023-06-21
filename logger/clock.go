package logger

import "time"

type constantClock time.Time

func (c constantClock) Now() time.Time { return time.Time(c) }
func (c constantClock) NewTicker(d time.Duration) *time.Ticker {
	return &time.Ticker{}
}
