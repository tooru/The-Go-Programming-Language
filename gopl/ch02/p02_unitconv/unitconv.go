package p02_unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }

const (
	OneFeet Meter = 0.3048
)

func MToF(m Meter) Feet { return Feet(m / OneFeet) }
func FToM(f Feet) Meter { return Meter(f) * OneFeet }

type Kirogram float64
type Pond float64

func (kg Kirogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (p Pond) String() string      { return fmt.Sprintf("%glb", p) }

const (
	OnePond Kirogram = 0.45359237
)

func KgToP(kg Kirogram) Pond { return Pond(kg / OnePond) }
func PToKg(p Pond) Kirogram  { return Kirogram(p) * OnePond }
