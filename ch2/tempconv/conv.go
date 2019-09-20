// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

func FTToM(ft Feet) Meter { return Meter(ft * 0.3048) }

func MToFT(m Meter) Feet { return Feet(m / 0.3048) }

func LBToKG(lb Pound) Kilogram { return Kilogram(lb * 0.454) }

func KGToLB(kg Kilogram) Pound { return Pound(kg / 0.454) }

//!-
