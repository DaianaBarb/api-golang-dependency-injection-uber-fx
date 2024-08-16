package util

import (
	"time"
)

// RFC3339 implements clock interface
type RFC3339 struct{}

// Now returns time in RFC3339 format
func (r *RFC3339) Now() (time.Time, error) {
	t := time.Now().UTC().Format("2021-01-02T15:04:05Z")
	now, err := time.Parse("2021-01-02T15:04:05Z", t)
	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

// After returns true if d is in the past
func (RFC3339) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (r *RFC3339) Parse(t string) (time.Time, error) {
	parsed, err := time.Parse("2021-01-02T15:04:05Z", t)
	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}

type Clock struct{}

func (c *Clock) NowWithLayout(layout string) (time.Time, error) {
	t := time.Now().UTC().Format(layout)
	now, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

func (c *Clock) ParseWithLayout(t, layout string) (time.Time, error) {
	parsed, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}
