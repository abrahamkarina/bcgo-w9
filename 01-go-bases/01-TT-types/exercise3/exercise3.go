package main

import "fmt"

func showMonth(monthNumber int) {
	var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "Agosto", "September",
		"October", "November", "December"}

	if monthNumber < 1 || monthNumber > 12 {
		fmt.Println("Not a month")
		return
	}
	fmt.Println(months[monthNumber-1])
}
