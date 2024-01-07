package lasagna_master

// PreparationTime recieves layers and the average layer prep time in minutes and returns the total average prep time.
func PreparationTime(layers []string, avgLayerPrepMins int) int {
	if avgLayerPrepMins == 0 {
		avgLayerPrepMins = 2
	}

	return (len(layers) * avgLayerPrepMins)
}

// Quantities takes the layers and returns the total noodles and sauce is needed to prepare them.
func Quantities(layers []string) (noodles int, sauce float64) {
	const litersSauce float64 = 0.2
	const gramsNoodles int = 50

	noodleLayers := 0
	sauceLayers := 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodleLayers++
		}
		if layers[i] == "sauce" {
			sauceLayers++
		}
	}

	noodles = gramsNoodles * noodleLayers
	sauce = litersSauce * float64(sauceLayers)
	return
}

// AddSecretIngredient replaces the last elemen ton customRecipes with the last element from recipes.
func AddSecretIngredient(recipes []string, customRecipes []string) {
	last := recipes[len(recipes)-1]
	customRecipes[len(customRecipes)-1] = last
}

// ScaleRecipe takes the desired ammounts for a pair of portions and the count and calculates the scaled ammount.
func ScaleRecipe(desiredAmmts []float64, desiredPortionCount int) []float64 {
	results := []float64{}
	for i := 0; i < len(desiredAmmts); i++ {
		sum := desiredAmmts[i] * float64(desiredPortionCount) / 2

		results = append(results, sum)
	}
	return results
}
