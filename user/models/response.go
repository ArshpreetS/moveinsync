package models

import "time"

type Response struct {
	TripID string `json:"tripid"`
	BusID  string `json:"busid"`
	Route  []struct {
		Stop     string    `json:"stop"`
		DateTime time.Time `json:"datetime"`
	} `json:"route"`
	Tickets   int       `json:"tickets"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
}
