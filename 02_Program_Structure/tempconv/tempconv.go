package tempconv

import "fmt"

// Celsius type is a temperature unit
type Celsius float64

// Fahrenheit type is a temperature unit
type Fahrenheit float64

const (
	// AbsoluteZeroC value in celcius degrees
	AbsoluteZeroC Celsius = -273.15
	// FreezingC value in celcius degrees
	FreezingC Celsius = 0
	// BoilingC value in celcius degrees
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
