// Package weather is used for getting weather forecast.
package weather

var (
	// CurrentCondition describes is a string that holds the current weather condition.
	CurrentCondition string
	// CurrentLocation describes is a string that holds the current location (Goblinocus).
	CurrentLocation string
)

// Forecast returns a weather current condition in a specified location.
// It waits city and condition as parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
