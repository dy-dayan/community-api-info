package util

import "strconv"

func Str2Int(p string) int {
	v, err := strconv.Atoi(p)
	if err != nil {
		return 0
	}
	return v
}

func Str2Int32(p string) int32 {
	v, err := strconv.Atoi(p)
	if err != nil {
		return 0
	}
	return int32(v)
}

func Str2Float32(p string) float32 {
	v, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return 0.0
	}
	return float32(v)
}

func Str2Float64(p string) float64 {
	v, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return 0.0
	}
	return v
}
