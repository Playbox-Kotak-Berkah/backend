package utils

func CalculateRealtime(pagi, siang, malam float64) float64 {
	result := 0.0
	divide := 0.0

	if pagi != 0 {
		result += pagi
		divide++
	}

	if siang != 0 {
		result += siang
		divide++
	}

	if malam != 0 {
		result += malam
		divide++
	}

	if result == 0.0 || divide == 0.0 {
		return 0
	}

	return result / divide
}
