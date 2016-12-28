package main

import (
	"fmt"
	"os"
	"strconv"
)

// temps
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// lengths
type Meter float64
type Foot float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (f Foot) String() string {
	if f > 1.0 || f < 1.0 {
		return fmt.Sprintf("%g feet", f)
	} else {
		return fmt.Sprintf("%g foot", f)
	}
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5/9)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func MetersToFeet(m Meter) Foot {
	return Foot(m * 3.28084)
}

func FeetToMeters(f Foot) Meter {
	return Meter(f * .3048)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uniter: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)

		feet := Foot(t)
		meters := Meter(t)
		fmt.Printf("Temps: %s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
		fmt.Printf("Lengths: %s = %s, %s = %s\n", feet, FeetToMeters(feet), meters, MetersToFeet(meters))
	}
}
