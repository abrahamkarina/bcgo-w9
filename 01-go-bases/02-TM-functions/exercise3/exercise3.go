package exercise3

import "errors"

type Category int

const (
	CategoryA = iota
	CategoryB
	CategoryC
)
const (
	salaryByHourCategoryA = 3000.0
	salaryByHourCategoryB = 1500.0
	salaryByHourCategoryC = 1000.0
)
const (
	bonificationCategoryA = .5
	bonificationCategoryB = .2
	bonificationCategoryC = .0
)

func minutesToHour(minutes float64) float64 {
	const minutesByHour = 60
	return float64(minutes) / minutesByHour
}
func CalculateSalary(minutesWorked float64, category Category) (float64, error) {
	worked := minutesToHour(minutesWorked)
	switch category {
	case CategoryA:
		return calculateSalary(worked, salaryByHourCategoryA, bonificationCategoryA), nil
	case CategoryB:
		return calculateSalary(worked, salaryByHourCategoryB, bonificationCategoryB), nil
	case CategoryC:
		return calculateSalary(worked, salaryByHourCategoryC, bonificationCategoryC), nil
	default:
		return 0, errors.New("invalid category")
	}
}

func calculateSalary(worked float64, salaryByHour float64, bonification float64) float64 {
	monthlySalary := worked * salaryByHour
	totalSalary := monthlySalary + monthlySalary*bonification
	return totalSalary
}
