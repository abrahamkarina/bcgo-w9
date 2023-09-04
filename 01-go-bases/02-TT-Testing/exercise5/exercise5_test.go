package exercise5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateFoodForCats(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		food := calculateFoodForCats(3)
		assert.Equal(t, 3*foodGramsForCat, food)
	})

	t.Run("ZeroInput", func(t *testing.T) {
		food := calculateFoodForCats(0)
		assert.Equal(t, 0.0, food)
	})
}

func TestCalculateFoodForDogs(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		food := calculateFoodForDogs(2)
		assert.Equal(t, 2*foodGramsForDog, food)
	})

	t.Run("ZeroInput", func(t *testing.T) {
		food := calculateFoodForDogs(0)
		assert.Equal(t, 0.0, food)
	})
}

func TestCalculateFoodForHamsters(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		food := calculateFoodForHamsters(5)
		assert.Equal(t, 5*foodGramsForHamster, food)
	})

	t.Run("ZeroInput", func(t *testing.T) {
		food := calculateFoodForHamsters(0)
		assert.Equal(t, 0.0, food)
	})
}

func TestCalculateFoodForTarantulas(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		food := calculateFoodForTarantulas(1)
		assert.Equal(t, foodGramsForTarantula, food)
	})

	t.Run("ZeroInput", func(t *testing.T) {
		food := calculateFoodForTarantulas(0)
		assert.Equal(t, 0.0, food)
	})
}

func TestAnimal(t *testing.T) {
	t.Run("ValidAnimalTypes", func(t *testing.T) {
		animalTypes := []struct {
			animalType       string
			expectedFunction func(int) float64
			expectedError    error
		}{
			{Dog, calculateFoodForDogs, nil},
			{Cat, calculateFoodForCats, nil},
			{Hamster, calculateFoodForHamsters, nil},
			{Tarantula, calculateFoodForTarantulas, nil},
		}

		for _, at := range animalTypes {
			fn, err := Animal(at.animalType)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("%p", at.expectedFunction), fmt.Sprintf("%p", fn))

		}
	})

	t.Run("InvalidAnimalType", func(t *testing.T) {
		_, err := Animal("Fish")
		require.Error(t, err)
		assert.EqualError(t, err, "invalid animal type")
	})
}
