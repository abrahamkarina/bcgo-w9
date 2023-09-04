package exercise3

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("CategoryA", func(t *testing.T) {
		salary, err := CalculateSalary(1200.0, CategoryA)
		require.NoError(t, err)
		assert.Equal(t, 90000.0, salary)
	})

	t.Run("CategoryB", func(t *testing.T) {
		salary, err := CalculateSalary(1200.0, CategoryB)
		require.NoError(t, err)
		assert.Equal(t, 36000.0, salary)
	})

	t.Run("CategoryC", func(t *testing.T) {
		salary, err := CalculateSalary(1200.0, CategoryC)
		require.NoError(t, err)
		assert.Equal(t, 20000.0, salary)
	})

	t.Run("InvalidCategory", func(t *testing.T) {
		_, err := CalculateSalary(1200.0, 3)
		require.Error(t, err)
		assert.EqualError(t, err, "invalid category")
	})
}

func TestCalculateSalaryInternal(t *testing.T) {
	t.Run("CategoryA", func(t *testing.T) {
		salary := calculateSalary(40.0, salaryByHourCategoryA, bonificationCategoryA)
		assert.Equal(t, 180000.0, salary)
	})

	t.Run("CategoryB", func(t *testing.T) {
		salary := calculateSalary(40.0, salaryByHourCategoryB, bonificationCategoryB)
		assert.Equal(t, 72000.0, salary)
	})

	t.Run("CategoryC", func(t *testing.T) {
		salary := calculateSalary(40.0, salaryByHourCategoryC, bonificationCategoryC)
		assert.Equal(t, 40000.0, salary)
	})
}
