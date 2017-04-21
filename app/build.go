package app

import (
	"strconv"
	"time"

	"trackstore/log"
)

// Version string -ldflags "-X trackstore/app.version=abcdefg"
var version string = "n/a"

// Returns the compiled version string
func Version() string {
	return version
}

// Time stamp string -ldflags "-X trackstore/app.timestamp=abcdefg"
var (
	timestamp        string
	defaultTimestamp time.Time = time.Now().UTC()
)

// Returns the compiled version string
func BuildTime() time.Time {
	if timestamp == "" {
		return defaultTimestamp
	}
	sec, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.WithError(err).Error("error parsing build timestamp")
		return defaultTimestamp
	}
	return time.Unix(sec, 0).UTC()
}
