// Package weather provide a function to print a city forecast.
package weather

var (
	// CurrentCondition represents a weather condition.
	CurrentCondition string
	// CurrentLocation represents a location.
	CurrentLocation string
)

// Forecast the city current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
