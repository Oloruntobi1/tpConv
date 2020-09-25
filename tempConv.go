package tpConv

// ToCelsius converts the given temperature
// in Fahreinheit to Celsius and returns the
// value
func ToCelsius(f float64) float64{
	return (f - 32) * 5 / 9
}

// ToFahreinheit converts the given temperature
// in Celsius to Fahreinheit and returns the
// value
func ToFahreinheit(c float64) float64{
	return (c * 9 / 5) + 32
}