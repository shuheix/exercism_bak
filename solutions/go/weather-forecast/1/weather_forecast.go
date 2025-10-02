// Package weather represents a current weather.
package weather

// CurrentCondition is a string value.
var CurrentCondition string
// CurrentLocation is a string value.
var CurrentLocation string


// Forecast combine CurrentLocation and CurrentCondition.
func Forecast (city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
