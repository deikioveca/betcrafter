package utils

import (
	"fmt"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", s)
	}

	*d = Date(t)
	return nil
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}