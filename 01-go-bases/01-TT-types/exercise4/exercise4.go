package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	printOlderThan(employees, 25)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}

func printOlderThan(employees map[string]int, age int) {
	for employeeName, employeeAge := range employees {
		if employeeAge > age {
			fmt.Println(employeeName)
		}
	}
}
