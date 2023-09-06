package main

import (
	"encoding/csv"
	"fmt"
	ticketP "github.com/abrahamkarina/bcgo-w9/01-go-bases/05-TT-repaso/desafio-go-bases/internal/tickets"
	"os"
	"strconv"
)

func main() {
	tickets, err := ReadTickets("tickets.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	destination := "China"

	totalTicketsToDestination, err := ticketP.GetTotalTickets(tickets, destination)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total of tickets to %s: %d\n", destination, totalTicketsToDestination)

	countByPeriod, err := ticketP.GetCountByPeriod(tickets, ticketP.Morning)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total of tickets for the morning: %d\n", countByPeriod)

	averageToDestination, err := ticketP.AverageDestination(tickets, destination)
	if err != nil {
		panic(err)

	}
	fmt.Printf("Average of tickets to %s destination: %.2f\n", destination, averageToDestination)
}

func ReadTickets(filename string) ([]ticketP.Ticket, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tickets []ticketP.Ticket
	for _, line := range lines {
		ticket := ticketP.Ticket{
			ID:          parseToInt(line[0]),
			Name:        line[1],
			Email:       line[2],
			Destination: line[3],
			Time:        line[4],
			Price:       parseToFloat(line[5]),
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func parseToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func parseToFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
