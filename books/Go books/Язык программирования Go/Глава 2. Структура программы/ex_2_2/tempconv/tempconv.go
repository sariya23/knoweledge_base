package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// AbsoluteZero - абсолютный ноль в градусах Цельсия.
	AbsoluteZero Celsius = -273.15
	// FreezingC - температура заморозки воды в градусах Цельсия.
	FreezingC Celsius = 0
	// BoilingC - температура закипания воды в градусах Цельсия.
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) Stirng() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

// CToF - переводит градусы Цельсия в градусы Фаренгейта.
func (c Celsius) CToF() Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK - переводит градусы Цельсия в градусы Кельвина.
func (c Celsius) CToK() Kelvin { return Kelvin(c + 273.15) }

// FToC - переводит градусы Фаренгейта в градусы Цельсия.
func (f Fahrenheit) FToC() Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK - переводит градусы Фаренгейта в градусы Кельвина.
func (f Fahrenheit) FToK() Kelvin { return Kelvin(f-32)*5/9 + 273.15 }

// KToC - переводит градусы Кельвина в градусы Цельсия.
func (k Kelvin) KToC() Celsius { return Celsius(k - 273.15) }

// KToF - переводит градусы Кельвина в градусы Фаренгейта.
func (k Kelvin) KToF() Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }
