package xmath

import "math"

const Epsilon = 0.0000001

func Operator(op string, a, b int) int {
	if op == "*" {
		return a * b
	} else if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "/" {
		return a / b
	} else if op == "%" {
		return a % b
	}
	return -1
}

func ToRomanNumeral(x int) string {
	switch {
	case x >= 1000:
		return "M" + ToRomanNumeral(x-1000)
	case x >= 900:
		return "CM" + ToRomanNumeral(x-900)
	case x >= 500:
		return "D" + ToRomanNumeral(x-500)
	case x >= 400:
		return "CD" + ToRomanNumeral(x-400)
	case x >= 100:
		return "C" + ToRomanNumeral(x-100)
	case x >= 90:
		return "XC" + ToRomanNumeral(x-90)
	case x >= 50:
		return "L" + ToRomanNumeral(x-50)
	case x >= 40:
		return "XL" + ToRomanNumeral(x-40)
	case x >= 10:
		return "X" + ToRomanNumeral(x-10)
	case x >= 9:
		return "IX" + ToRomanNumeral(x-9)
	case x >= 5:
		return "V" + ToRomanNumeral(x-5)
	case x >= 4:
		return "IV" + ToRomanNumeral(x-4)
	case x >= 1:
		return "I" + ToRomanNumeral(x-1)
	}
	return ""
}

func Round(x float64) float64 {
	return ToNearestEven(x)
}

func RoundTo(x float64, dp float64) float64 {
	x = x * math.Pow(10, dp)
	return ToNearestEven(x) / math.Pow(10, dp)
}

func ToNearestEven(x float64) float64 {
	return toNearest(x, true)
}

func ToNearestAway(x float64) float64 {
	return toNearest(x, false)
}

func ToZero(x float64) float64 {
	return math.Trunc(x)
}

func AwayFromZero(x float64) float64 {
	if x >= 0 {
		return math.Ceil(x)
	} else {
		return math.Floor(x)
	}
}

func ToPositiveInf(x float64) float64 {
	return math.Ceil(x)
}

func ToNegativeInf(x float64) float64 {
	return math.Floor(x)
}

func toNearest(x float64, tiesToEven bool) float64 {
	if x == 0 || math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}
	if x < 0.0 {
		return -toNearest(-x, tiesToEven)
	}

	intPart, fracPart := math.Modf(x)

	if math.Abs(fracPart-0.5) < Epsilon {
		if tiesToEven {
			if math.Mod(intPart, 2.0) < Epsilon {
				return intPart
			}
		}
		return math.Ceil(intPart + 0.5)
	}
	return math.Floor(x + 0.5)
}
