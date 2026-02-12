// Package weather provides tools that return the current weather condition of a city.
package weather

// CurrentCondition represents the current weather condition of a city.
var CurrentCondition string

// CurrentLocation represents the city for which the current weather condition is reported.
var CurrentLocation string

// Forecast returns a string that describes the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
