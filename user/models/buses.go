package models

import "time"

type Bus_data struct {
	BusID   string `json:"busid"`
	Seats   int    `json:"seats"`
	BusName string `json:"busName"`
}

type Ticket struct {
	TripID    string    `json:"tripid"`
	BusID     string    `json:"BusID"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
}
