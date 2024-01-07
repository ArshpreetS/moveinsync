package utils

import (
	"github.com/ArshpreetS/moveinsync/user/models"
)

func FilterZeroTickets(v []models.Response) []models.Response {
	filtered := []models.Response{}

	for _, val := range v {
		if val.Tickets != 0 {
			filtered = append(filtered, val)
		}
	}

	return filtered
}
