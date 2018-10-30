package state_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete/state"
	"github.com/kzahedi/gomi/discrete"
)

func TestConditionalMutualInformationBaseE(t *testing.T) {
	t.Log("Testing Conditional Mutual Information")
	data := make([][]int, 100)
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			data[i] = make([]int, 3, 3)
			data[i][0] = int(rand.Int63n(100))
			data[i][1] = int(rand.Int63n(100))
			data[i][2] = int(rand.Int63n(100))
		}

		r := state.ConditionalMutualInformationBaseE(data)
		p := discrete.Empirical3D(data)
		mi := discrete.ConditionalMutualInformationBaseE(p)

		s := 0.0
		for i := 0; i < len(r); i++ {
			s += r[i]
		}

		s /= float64(len(r))

		if math.Abs(mi-s) > 0.00001 {
			t.Errorf("Conditional Mutual information should be equal %f = %f", s, mi)
		}
	}
}

func TestConditionalMutualInformationBase2(t *testing.T) {
	t.Log("Testing Conditional Mutual Information")
	data := make([][]int, 100)
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			data[i] = make([]int, 3, 3)
			data[i][0] = int(rand.Int63n(100))
			data[i][1] = int(rand.Int63n(100))
			data[i][2] = int(rand.Int63n(100))
		}

		r := ConditionalMutualInformationBase2(data)
		p := discrete.Empirical3D(data)
		mi := discrete.ConditionalMutualInformationBase2(p)

		s := 0.0
		for i := 0; i < len(r); i++ {
			s += r[i]
		}

		s /= float64(len(r))

		if math.Abs(mi-s) > 0.00001 {
			t.Errorf("Conditional Mutual information should be equal %f = %f", s, mi)
		}
	}
}
