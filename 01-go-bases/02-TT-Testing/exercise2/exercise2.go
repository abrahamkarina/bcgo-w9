package exercise2

import "errors"

func CalculateAverage(values ...float64) (float64, error) {
	n := len(values)
	if n == 0 {
		return 0, nil
	}
	count := .0
	for _, val := range values {
		if val < 0 {
			return 0, errors.New("negative value")
		}
		count += val
	}
	return count / float64(n), nil

}
