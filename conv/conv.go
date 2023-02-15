package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (value - 32) * 5 / 9
}

// De andre konverteringsfunksjonene implementere her
// ...

// Konverterer Celcius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {

	return value*9/5 + 32
}

// Konverterer Kelvin til Celcius
func KelvinToCelcius(value float64) float64 {

	return value - 273.15
}

// Konverterer Celcius til Kelvin
func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {

	return (value-273.15)*9/5 + 32
}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {

	return (value-32)*5/9 + 273.15
}
