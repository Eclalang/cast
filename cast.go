package cast

import "strconv"

// Atoi converts a string to an integer
func Atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// FloatToInt converts a float to an integer
func FloatToInt(f float64) int {
	return int(f)
}

// IntToFloat converts an integer to a float
func IntToFloat(i int) float64 {
	return float64(i)
}

// ParseBool converts a string to a boolean
func ParseBool(str string) bool {
	boolean, _ := strconv.ParseBool(str)
	return boolean
}

// ParseFloat converts a string to a float
func ParseFloat(str string, bitSize int) float64 {
	f, _ := strconv.ParseFloat(str, bitSize)
	return f
}
