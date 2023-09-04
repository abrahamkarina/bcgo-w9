package exercise1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTaxes(t *testing.T) {
	testCases := []struct {
		salary           float64
		expectedTax      float64
		expectedErrorMsg string
	}{
		{
			salary:      40000,
			expectedTax: 0,
		},
		{
			salary:      70000,
			expectedTax: 70000 * 0.17,
		},
		{
			salary:      200000,
			expectedTax: 200000 * 0.27,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Salary: %.2f", tc.salary), func(t *testing.T) {
			result := CalculateTaxes(tc.salary)
			assert.Equal(t, tc.expectedTax, result, "Tax calculation does not match expected result")
		})
	}
}
