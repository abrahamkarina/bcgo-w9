package exercise1

const (
	firstCategory  = 50000
	secondCategory = 150000
)
const (
	taxFirstCategory  = .17
	taxSecondCategory = .27
)

func CalculateTaxes(salary float64) float64 {
	if salary >= secondCategory {
		return salary * taxSecondCategory
	}
	if salary >= firstCategory {
		return salary * taxFirstCategory
	}
	return 0
}
