package tickets

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrZeroTickets = errors.New("cannot calculate average, at least a ticket is needed")
)

const (
	Dawn = iota
	Morning
	Afternoon
	Night
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       float64
}


func GetTotalTickets(tickets []Ticket, destination string) (int, error) {
	count := 0
	for _, ticket := range tickets {
		if strings.EqualFold(ticket.Destination, destination) {
			count++
		}
	}
	return count, nil
}

func GetCountByPeriod(tickets []Ticket, period int) (int, error) {
	count := 0
	for _, ticket := range tickets {
		t, err := time.Parse("15:04", ticket.Time)
		if err != nil {
			return 0, err
		}

		switch period {
		case Dawn:
			if t.Hour() >= 0 && t.Hour() < 6 {
				count++
			}
		case Morning:
			if t.Hour() >= 7 && t.Hour() < 12 {
				count++
			}
		case Afternoon:
			if t.Hour() >= 13 && t.Hour() < 19 {
				count++
			}
		case Night:
			if t.Hour() >= 20 && t.Hour() <= 23 {
				count++
			}
		}
	}
	return count, nil
}

func AverageDestination(tickets []Ticket, destination string) (float64, error) {
	total := len(tickets)
	if total == 0 {
		return 0, ErrZeroTickets
	}
	count, err := GetTotalTickets(tickets, destination)
	if err != nil {
		return 0, err
	}

	return float64(count) / float64(total), nil
}
