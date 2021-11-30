// Package weather provides tools for Wheather conditions.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation describes the current location.
var CurrentLocation string

// Forecast function returns the conditions for a location and a weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
