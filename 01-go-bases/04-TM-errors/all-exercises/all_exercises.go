package main

import (
	"errors"
	"fmt"
)

var (
	NoAlcanzaMinimoImponible = NewMinimoNoImponibleError()
	HoursError               = errors.New("el trabajador no puede haber trabajado menos de  80 hs mensuales")
)

const (
	minimoNoImponible = 150000
)

// CalculateTaxes1and2 CalculateTaxes for exercises 1 and 2
func CalculateTaxes1and2(salary float64) error {
	if salary < minimoNoImponible {
		return NoAlcanzaMinimoImponible
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

// CalculateTaxes3 CalculateTaxes for exercise 3
func CalculateTaxes3(salary float64) error {
	if salary < minimoNoImponible {
		return errors.New(NoAlcanzaMinimoImponible.Error())
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

// CalculateTaxes4 CalculateTaxes for exercise 4
func CalculateTaxes4(salary float64) error {
	if salary < minimoNoImponible {
		//return fmt.Errorf("error: %w el salario ingresado es de %f", NoAlcanzaMinimoImponible, salary)
		return fmt.Errorf("error: el minimo no imponible es de 150000 y el salario ingresado es de %f", salary)
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

// CalculateTaxes5 CalculateTaxes for exercise 5
func CalculateTaxes5(salary float64) (float64, error) {
	const tax = 0.1
	if salary < minimoNoImponible {
		return 0, NoAlcanzaMinimoImponible

	}

	return salary * tax, nil
}

// CalculateSalary for exercise 5
func CalculateSalary(salaryByHour, workedHours float64) (float64, error) {
	const minHours = 80
	if workedHours < minHours {
		return 0, HoursError
	}
	salary := salaryByHour * workedHours
	tax, err := CalculateTaxes5(salary)
	if err != nil {
		if errors.Is(err, NoAlcanzaMinimoImponible) {
			return salary, nil
		} else {
			return 0, err
		}
	}
	return salary - tax, nil
}
func main() {
	salary := 10000.0

	//exercises 1 and 2
	err := CalculateTaxes1and2(salary)
	if errors.Is(err, NoAlcanzaMinimoImponible) {
		fmt.Println("2)", err.Error())
	}

	//exercise 3
	err = CalculateTaxes3(salary)
	if !errors.Is(err, NoAlcanzaMinimoImponible) {
		fmt.Println("3)", err.Error())
	}

	//exercise 4
	err = CalculateTaxes4(salary)
	if !errors.Is(err, NoAlcanzaMinimoImponible) {
		fmt.Println("4)", err.Error())
	}

	//exercise 5
	_, err = CalculateSalary(100, 40)
	if err != nil {
		fmt.Println("5)a)", err.Error())
	}
	salaryExercise5, err := CalculateSalary(10000, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println("5)b)", salaryExercise5)

	salaryExercise5, err = CalculateSalary(100, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println("5)c)", salaryExercise5)
}
