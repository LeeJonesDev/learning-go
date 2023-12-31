// Package weather forecasts the current location's weather.
package weather

// CurrentCondition of the weather.
var CurrentCondition string

// CurrentLocation of a user seeking a weather forcast.
var CurrentLocation string

// Forecast takes the current city and current weather conditions as strings and returns the forecast as a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
