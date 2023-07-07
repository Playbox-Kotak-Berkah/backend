package utils

import (
	"strconv"
)

func StringToInteger(s string) int {
	converted, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return converted
}

func StringToFloat64(s string) float64 {
	converted, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return converted
}

func StringToBool(s string) bool {
	converted, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return converted
}
