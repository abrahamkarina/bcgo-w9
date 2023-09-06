package tickets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	tickets := []Ticket{
		Ticket{Destination: "Argentina"},
		Ticket{Destination: "Uruguay"},
		Ticket{Destination: "Chile"},
		Ticket{Destination: "Argentina"},
	}

	tests := []struct {
		destination     string
		expectedTotal   int
		expectedError   error
		testDescription string
	}{
		{"Argentina", 2, nil, "Two tickets to Argentina"},
		{"China", 0, nil, "No tickets to China"},
	}

	for _, test := range tests {
		t.Run(test.testDescription, func(t *testing.T) {
			total, err := GetTotalTickets(tickets, test.destination)
			assert.Equal(t, test.expectedTotal, total)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
func TestGetCountByPeriod(t *testing.T) {
	tickets := []Ticket{
		Ticket{Time: "08:30"},
		Ticket{Time: "15:45"},
		Ticket{Time: "20:15"},
		Ticket{Time: "02:30"},
	}

	tests := []struct {
		period          int
		expectedCount   int
		expectedError   error
		testDescription string
	}{
		{Dawn, 1, nil, "dawn"},
		{Morning, 1, nil, "morning"},
		{Afternoon, 1, nil, "afternoon"},
		{Night, 1, nil, "night"},
	}

	for _, test := range tests {
		t.Run(test.testDescription, func(t *testing.T) {
			count, err := GetCountByPeriod(tickets, test.period)
			assert.Equal(t, test.expectedCount, count)
			assert.Equal(t, test.expectedError, err)
		})
	}

}

func TestAverageDestination(t *testing.T) {
	tickets := []Ticket{
		Ticket{Destination: "Argentina"},
		Ticket{Destination: "Uruguay"},
		Ticket{Destination: "Chile"},
		Ticket{Destination: "Argentina"},
	}

	tests := []struct {
		tickets         []Ticket
		destination     string
		expectedError   error
		expectedAverage float64
		testDescription string
	}{
		{
			tickets:         tickets,
			destination:     "North Korea",
			expectedError:   nil,
			expectedAverage: 0,
			testDescription: "No tickets to north korea",
		},
		{
			tickets:         tickets,
			destination:     "Argentina",
			expectedError:   nil,
			expectedAverage: 0.5,
			testDescription: "Average of 0.5 tickets to Argentina",
		},
		{
			tickets:         nil,
			destination:     "Argentina",
			expectedError:   ErrZeroTickets,
			expectedAverage: 0,
			testDescription: "No tickets",
		},
	}
	for _, test := range tests {
		t.Run(test.testDescription, func(t *testing.T) {
			average, err := AverageDestination(test.tickets, test.destination)
			assert.Equal(t, test.expectedAverage, average)
			assert.Equal(t, test.expectedError, err)
		})
	}

}
