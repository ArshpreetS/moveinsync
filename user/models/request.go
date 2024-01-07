package models

import "time"

type TravelData struct {
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Date        time.Time `json:"date"`
}

type TicketToBook struct {
	TripID    string `json:"tripid"`
	BusID     string `json:"busid"`
	From      string `json:"from"`
	To        string `json:"To"`
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
}
