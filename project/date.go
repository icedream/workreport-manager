package project

import (
	"fmt"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

// Date is a wrapper around time.Time with a different UnmarshalJSON
// implementation that parses just date information.
type Date struct {
	time.Time
}

// UnmarshalJSON takes the given buffer and translates it into a proper
// time.Time object.
func (t *Date) UnmarshalJSON(buf []byte) (err error) {
	tt, err := time.Parse(dateFormat, strings.Trim(string(buf), `"`))
	if err != nil {
		return
	}
	t.Time = tt
	return nil
}

func (t Date) String() string {
	// TODO - Localize date!
	return fmt.Sprintf("%02d.%02d.%04d", t.Time.Day(), t.Time.Month(), t.Time.Year())
}
