package models

import (
	"time"
)

type Bus_data struct {
	BusID   string `json:"busid"`
	Seats   int    `json:"seats"`
	BusName string `json:"busName"`
}

type Bus_Trip struct {
	BusID  string `json:"busid"`
	TripID string `json:"tripid"`
	Route  []struct {
		Stop     string    `json:"stop"`
		DateTime time.Time `json:"datetime"`
	} `json:"route"`

	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
}
