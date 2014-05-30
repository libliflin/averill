package averill

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	samples := []float64{17.72, 19.75, 24.91, 19.65, 27.80}

	for i := 0; i < len(samples); i++ {
		fmt.Println(samples[i], Compound_from_mL_and_g(25.00, samples[i]))
	}

}
