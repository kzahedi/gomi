package state

import (
	"github.com/kzahedi/goent/continuous/state"
)

func diff(r1, r2 []float64) []float64 {
	r := make([]float64, len(r1), len(r1))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationW [...]
func MorphologicalComputationW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	return state.FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
}

// MorphologicalComputationA [...]
func MorphologicalComputationA(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	return state.FrenzelPompe(w2w1a1, w2Indices, a1Indices, w1Indices, k, eta)
}

// MorphologicalComputationMI1 [...]
func MorphologicalComputationMI1(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) []float64 {
	r1 := state.KraskovStoegbauerGrassberger1(w2w1s1a1, w2Indices, w1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger1(w2w1s1a1, s1Indices, a1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationMI2 [...]
func MorphologicalComputationMI2(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) []float64 {
	r1 := state.KraskovStoegbauerGrassberger2(w2w1s1a1, w2Indices, w1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger2(w2w1s1a1, a1Indices, s1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationCA1 quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCA = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	r1 := state.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationCA2 quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCA = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCA2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	r1 := state.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationWA1 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	n := len(w1Indices) + len(a1Indices)
	w1a1Indices := make([]int, n, n)
	index := 0
	for _, v := range w1Indices {
		w1a1Indices[index] = v
		index++
	}
	for _, v := range a1Indices {
		w1a1Indices[index] = v
		index++
	}

	r1 := state.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1a1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationWA2 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) []float64 {
	n := len(w1Indices) + len(a1Indices)
	w1a1Indices := make([]int, n, n)
	index := 0
	for _, v := range w1Indices {
		w1a1Indices[index] = v
		index++
	}
	for _, v := range a1Indices {
		w1a1Indices[index] = v
		index++
	}

	r1 := state.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1a1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationWS1 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS1(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) []float64 {
	n := len(w1Indices) + len(s1Indices)
	w1s1Indices := make([]int, n, n)
	index := 0
	for _, v := range w1Indices {
		w1s1Indices[index] = v
		index++
	}
	for _, v := range s1Indices {
		w1s1Indices[index] = v
		index++
	}

	r1 := state.KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, w1s1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, s1Indices, k, eta)
	return diff(r1, r2)
}

// MorphologicalComputationWS2 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS2(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) []float64 {
	n := len(w1Indices) + len(s1Indices)
	w1s1Indices := make([]int, n, n)
	index := 0
	for _, v := range w1Indices {
		w1s1Indices[index] = v
		index++
	}
	for _, v := range s1Indices {
		w1s1Indices[index] = v
		index++
	}

	r1 := state.KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, w1s1Indices, k, eta)
	r2 := state.KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, s1Indices, k, eta)
	return diff(r1, r2)
}
