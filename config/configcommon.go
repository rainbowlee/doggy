package config

import (
	"strconv"
)

func getIntValue(data string) int {
	intvalue, error := strconv.Atoi(data)
	if error == nil {
		return intvalue;
	}
	return 0
}

func getInt64Value(data string) int64 {
	int64value, error := strconv.ParseInt(data, 10, 64)
	if error == nil {
		return int64value;
	}
	return 0
}

func getFloat64Value(data string) float64 {
	float64value, error := strconv.ParseFloat(data, 64)
	if error == nil {
		return float64value;
	}
	return 0
}

func getFloat32Value(data string) float32 {
	float32value, error := strconv.ParseFloat(data, 32)
	if error == nil {
		return float32(float32value);
	}	
	return 0
}