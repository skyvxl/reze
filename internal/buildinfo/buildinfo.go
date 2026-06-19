package buildinfo

import (
	"fmt"
	"time"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func String() string {
	date := BuildDate
	if t, err := time.Parse(time.RFC3339, date); err == nil {
		date = t.Format("2006-01-02")
	}
	return fmt.Sprintf("%s-%s-%s", Version, Commit, date)
}
