package sts

func SomaFloat64(x []float64) float64 {
	total := 0.

	for _, v := range x {
		total += v
	}

	return total
}

func AvgFloat64(x []float64) float64 {
	total := SomaFloat64(x)

	return float64(total) / float64(len(x))
}
