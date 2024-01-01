package cars

const CarCost int = 10000
const DiscountCarCost int = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCarsPerHour := (float64(productionRate) * successRate)
	workingCarsPerHourAsPercentage := workingCarsPerHour * .01

	return workingCarsPerHourAsPercentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	// modulo 1 == 1 and should be full price
	if carsCount == 1 {
		return uint(CarCost)
	}

	// get remainder divided by 10 for full price cars
	fullPriceUnits := carsCount % 10

	// the count of the rest is the count less the full price units
	discountedUnits := carsCount - fullPriceUnits

	// discounted units number divided by 10 * discounted car cost cost ex: 30 == 3 * 95000
	discountedUnitsCost := float64(discountedUnits/10) * float64(DiscountCarCost)

	fullPriceUnitsCost := fullPriceUnits * CarCost

	return uint(discountedUnitsCost + float64(fullPriceUnitsCost))
}
