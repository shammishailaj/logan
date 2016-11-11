package filter

import (
	"time"

	"github.com/kbence/logan/types"
)

type TimeFilter struct {
	Start time.Time
	End   time.Time
}

func NewTimeFilter(start time.Time, end time.Time) *TimeFilter {
	return &TimeFilter{Start: start, End: end}
}

func (f *TimeFilter) Match(line *types.LogLine) bool {
	return (line.Date.After(f.Start) || line.Date.Equal(f.Start)) &&
		(line.Date.Before(f.End) || line.Date.Equal(f.End))
}
