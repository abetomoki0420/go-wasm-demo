package math

func AddCore(a float64, b float64) float64 {
	return a + b
}

func SumCore(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return float64(sum)
}
