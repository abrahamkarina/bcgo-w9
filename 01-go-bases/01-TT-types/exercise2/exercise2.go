package main

import "fmt"

const (
	minAge                = 22
	minYearsEmployed      = 1
	salaryWithoutInterest = 100.00
)

func loan(age int, isEmployed bool, yearsEmployed float64, salary float64) {
	if age < minAge || !isEmployed || yearsEmployed < minYearsEmployed {
		fmt.Println("Load denied")
		return
	}
	fmt.Println("Load accepted")
	if salary >= salaryWithoutInterest {
		fmt.Println("No interest will be charged")
	}
}
