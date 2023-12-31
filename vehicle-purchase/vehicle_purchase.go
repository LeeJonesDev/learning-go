package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	appendMessage := " is clearly the better choice."

	if option1 < option2 {
		return option1 + appendMessage
	}
	return option2 + appendMessage
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return .8 * originalPrice
	} else if age >= 3 && age < 10 {
		return .7 * originalPrice
	}
	return .5 * originalPrice

}
