package clock

import "fmt"

const minutesPerDay = 24 * 60

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := (h*60 + m) % minutesPerDay
	if minutes < 0 {
		minutes += minutesPerDay
	}
	return Clock{minutes}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
