package exercise4

import (
	"errors"
	"sort"
)

const (
	Minimum = iota
	Average
	Maximum
)

type OperationFunc func(values ...int) (float64, error)

func operation(operationType int) (OperationFunc, error) {

	switch operationType {
	case Minimum:
		return calculateMin, nil
	case Average:
		return calculateAverage, nil
	case Maximum:
		return calculateMax, nil
	default:
		return nil, errors.New("not defined operation")

	}
}

func calculateMin(values ...int) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("not enough values")
	}
	sort.Ints(values)
	return float64(values[0]), nil

}
func calculateAverage(values ...int) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("not enough values")
	}
	count := 0
	for _, val := range values {
		count += val
	}
	return float64(count) / float64(len(values)), nil
}
func calculateMax(values ...int) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("not enough values")
	}
	sort.Ints(values)
	return float64(values[len(values)-1]), nil
}
