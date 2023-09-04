package exercise4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateMin(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		min, err := calculateMin()
		require.Error(t, err)
		assert.EqualError(t, err, "not enough values")
		assert.Equal(t, float64(0), min)
	})

	t.Run("SingleValue", func(t *testing.T) {
		min, err := calculateMin(5)
		require.NoError(t, err)
		assert.Equal(t, float64(5), min)
	})

	t.Run("MultipleValues", func(t *testing.T) {
		min, err := calculateMin(10, 2, 8, 1, 5)
		require.NoError(t, err)          // No error should occur
		assert.Equal(t, float64(1), min) // Minimum should be 1
	})
}

func TestCalculateAverage(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		avg, err := calculateAverage()
		require.Error(t, err)                          // Error should occur for empty input
		assert.EqualError(t, err, "not enough values") // Error message should match
		assert.Equal(t, float64(0), avg)               // Average should be 0
	})

	t.Run("SingleValue", func(t *testing.T) {
		avg, err := calculateAverage(5)
		require.NoError(t, err) // No error should occur
		assert.Equal(t, float64(5), avg)
	})

	t.Run("MultipleValues", func(t *testing.T) {
		avg, err := calculateAverage(10, 2, 8, 1, 5)
		require.NoError(t, err)
		assert.Equal(t, 5.2, avg)
	})
}

func TestCalculateMax(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		max, err := calculateMax()
		require.Error(t, err)                          // Error should occur for empty input
		assert.EqualError(t, err, "not enough values") // Error message should match
		assert.Equal(t, float64(0), max)               // Maximum should be 0
	})

	t.Run("SingleValue", func(t *testing.T) {
		max, err := calculateMax(5)
		require.NoError(t, err)          // No error should occur
		assert.Equal(t, float64(5), max) // Maximum should be the single value
	})

	t.Run("MultipleValues", func(t *testing.T) {
		max, err := calculateMax(10, 2, 8, 1, 5)
		require.NoError(t, err)           // No error should occur
		assert.Equal(t, float64(10), max) // Maximum should be 10
	})
}

func TestOperation(t *testing.T) {
	t.Run("ValidOperations", func(t *testing.T) {
		operations := []struct {
			operationType int
			expectedFunc  OperationFunc
			expectedError error
		}{
			{Minimum, calculateMin, nil},
			{Average, calculateAverage, nil},
			{Maximum, calculateMax, nil},
		}

		for _, op := range operations {
			fn, err := operation(op.operationType)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("%p", op.expectedFunc), fmt.Sprintf("%p", fn))
			assert.NoError(t, err)
		}
	})

	t.Run("InvalidOperation", func(t *testing.T) {
		_, err := operation(42)                            // Using an invalid operation type
		require.Error(t, err)                              // An error should occur
		assert.EqualError(t, err, "not defined operation") // Error message should match
	})
}
