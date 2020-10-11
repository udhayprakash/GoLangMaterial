package main

import (
	"fmt"
	"time"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {
	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"America/Metropolis",
		"Asia/Kolkata",
		"Africa/Lagos",
		"Asia/karachi",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Format("2006-01-02 15:04"), t.Location())
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}
