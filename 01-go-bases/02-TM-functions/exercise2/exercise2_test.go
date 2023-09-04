package exercise2

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		average, err := CalculateAverage()
		require.NoError(t, err)
		assert.Equal(t, float64(0), average)
	})

	t.Run("PositiveValues", func(t *testing.T) {
		average, err := CalculateAverage(10.0, 20.0, 30.0)
		require.NoError(t, err)
		assert.Equal(t, float64(20), average)
	})

	t.Run("NegativeValueError", func(t *testing.T) {
		average, err := CalculateAverage(10.0, -5.0, 30.0)
		require.Error(t, err)
		assert.Equal(t, float64(0), average)
		assert.EqualError(t, err, "negative value")
	})
}
