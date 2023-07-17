// Package weather forecast the current weather condition of cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition of a city.
var CurrentCondition string
// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast returns the current weather condition of a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
