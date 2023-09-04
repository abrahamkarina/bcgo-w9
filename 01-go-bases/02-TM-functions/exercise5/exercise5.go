package exercise5

import "errors"

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)
const (
	foodGramsForDog       = 10000.0
	foodGramsForCat       = 5000.0
	foodGramsForHamster   = 250.0
	foodGramsForTarantula = 150.0
)

func Animal(animalType string) (func(int) float64, error) {
	switch animalType {
	case Dog:
		return calculateFoodForDogs, nil
	case Cat:
		return calculateFoodForCats, nil
	case Hamster:
		return calculateFoodForHamsters, nil
	case Tarantula:
		return calculateFoodForTarantulas, nil
	default:
		return nil, errors.New("invalid animal type")
	}
}

func calculateFood(animalNumber int, gramForUnit float64) float64 {
	return float64(animalNumber) * gramForUnit
}
func calculateFoodForCats(animalNumber int) float64 {
	return float64(animalNumber) * foodGramsForCat

}

func calculateFoodForDogs(animalNumber int) float64 {
	return float64(animalNumber) * foodGramsForDog

}

func calculateFoodForHamsters(animalNumber int) float64 {
	return float64(animalNumber) * foodGramsForHamster

}

func calculateFoodForTarantulas(animalNumber int) float64 {
	return float64(animalNumber) * foodGramsForTarantula

}
