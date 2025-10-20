package lasagna

import "fmt"

var OvenTime = 40

func RemainingOvenTime(actual int) int {
	var actualMinutesInOven int = OvenTime - actual
	return actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) {
	var Time int = numberOfLayers + actualMinutesInOven
	fmt.Println(Time)
}

func main() {
	var actual int = 30
	var numberOfLayers int = 5

	var actualMinutesInOven int = RemainingOvenTime(actual)
	numberOfLayers = PreparationTime(numberOfLayers)
	ElapsedTime(numberOfLayers, actualMinutesInOven)
}
