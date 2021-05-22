// Simple unit conversion program.
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Pounds float64
type Kilograms float64
type Feet float64
type Meters float64

const (
	KgPoundConversion    = 2.205
	MetersFeetConversion = 3.281
)

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func PoundToKg(p Pounds) Kilograms {
	return Kilograms(p / KgPoundConversion)
}

func KgToPound(k Kilograms) Pounds {
	return Pounds(k * KgPoundConversion)
}

func FtToM(f Feet) Meters {
	return Meters(f / MetersFeetConversion)
}

func MToFt(m Meters) Feet {
	return Feet(m * MetersFeetConversion)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g lbs", p)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%g kgs", k)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m", m)
}

func printConversions(t float64) {
	f := Fahrenheit(t)
	c := Celsius(t)
	lbs := Pounds(t)
	kgs := Kilograms(t)
	ft := Feet(t)
	m := Meters(t)

	fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", lbs, PoundToKg(lbs), kgs, KgToPound(kgs))
	fmt.Printf("%s = %s, %s = %s\n", ft, FtToM(ft), m, MToFt(m))
	fmt.Println("")
}

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		printConversions(t)
	}

}
