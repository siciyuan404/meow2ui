package cost

import (
	"math"
	"testing"
)

func TestEstimateCost(t *testing.T) {
	tests := []struct {
		name     string
		inTok    int
		outTok   int
		inPer1k  float64
		outPer1k float64
		expected float64
	}{
		{"zero tokens", 0, 0, 0.01, 0.03, 0.0},
		{"1k input only", 1000, 0, 0.01, 0.03, 0.01},
		{"1k output only", 0, 1000, 0.01, 0.03, 0.03},
		{"mixed", 500, 200, 0.01, 0.03, 0.005 + 0.006},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateCost(tt.inTok, tt.outTok, tt.inPer1k, tt.outPer1k)
			if math.Abs(got-tt.expected) > 0.0001 {
				t.Errorf("EstimateCost() = %f, want %f", got, tt.expected)
			}
		})
	}
}
