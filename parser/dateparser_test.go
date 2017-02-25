package parser

import (
	"testing"
	"time"
)

func ExpectParsedDate(t *testing.T, line string, date time.Time) {
	parsedDate := ParseDate(line)

	if parsedDate == nil {
		t.Errorf("Date from line '%s' couldn't be parsed (ParseDate returned nil)!", line)
		return
	}

	if *parsedDate != date {
		t.Errorf("Parsed date from line '%s' (parsed as '%s') doesn't match expected '%s'!",
			line, *parsedDate, date)
	}
}

func TestParseDateFindsDateInLog(t *testing.T) {
	previousLocationName := defaultLocationName
	defaultLocationName = "UTC"

	now := time.Now()
	exactDate := time.Date(2016, 12, 5, 6, 57, 36, 0, getLocation())
	date := time.Date(now.Year(), 12, 5, 6, 57, 36, 0, getLocation())
	yearEnd := time.Date(now.Year()-1, 12, 31, 23, 59, 58, 0, getLocation())

	if now.Before(date) {
		date = time.Date(date.Year()-1, 12, 5, 6, 57, 36, 0, getLocation())
	}

	ExpectParsedDate(t, "2016-12-05 06:57:36,000 This is a test log line", exactDate)
	ExpectParsedDate(t, "2016-12-05 06:57:36.000 This is a test log line", exactDate)
	ExpectParsedDate(t, "2016-12-05 06:57:36,000+0100 This is a test log line", exactDate)
	ExpectParsedDate(t, "2016-12-05 06:57:36,000-0100 This is a test log line", exactDate)
	ExpectParsedDate(t, "Dec  5 06:57:36 2016 This is a test log line", date)
	ExpectParsedDate(t, "Dec 5 06:57:36 2016 This is a test log line", date)
	ExpectParsedDate(t, "Dec  5 06:57:36 This is a test log line", date)
	ExpectParsedDate(t, "Dec 5 06:57:36 This is a test log line", date)
	ExpectParsedDate(t, "Mon, 05 Dec 06:57:36 UTC This is a test log line", date)
	ExpectParsedDate(t, "Mon, 05 Dec 06:57:36.000 +0000 This is a test log line", date)
	ExpectParsedDate(t, "Mon, 05 Dec 06:57:36.000 UTC This is a test log line", date)
	ExpectParsedDate(t, "Mon 05 Dec 06:57:36 UTC This is a test log line", date)
	// This test might fail for two minutes per year:
	ExpectParsedDate(t, "Dec 31 23:59:58 This is a test log line", yearEnd)

	defaultLocationName = previousLocationName
}
