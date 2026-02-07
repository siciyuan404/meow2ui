package cost

func EstimateCost(inputTokens int, outputTokens int, inputPer1k float64, outputPer1k float64) float64 {
	return (float64(inputTokens)/1000.0)*inputPer1k + (float64(outputTokens)/1000.0)*outputPer1k
}
