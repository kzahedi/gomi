package continuous

import (
	"github.com/kzahedi/goent/continuous"
)

// MorphologicalComputationW [...]
func MorphologicalComputationW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
}

// MorphologicalComputationA [...]
func MorphologicalComputationA(w2w1a1 [][]float64, w2Indices, a1Indices, w1Indices []int, k int, eta bool) float64 {
	return continuous.FrenzelPompe(w2w1a1, w2Indices, a1Indices, w1Indices, k, eta)
}

// MorphologicalComputationCW1 [...]
func MorphologicalComputationCW1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, false) - continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationCW2 [...]
func MorphologicalComputationCW2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, false) - continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWA1 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
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

	return continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1a1Indices, k, false) - continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWA2 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
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

	return continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1a1Indices, k, false) - continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWS1 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS1(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) float64 {
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

	return continuous.KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, w1s1Indices, k, false) - continuous.KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, s1Indices, k, eta)
}

// MorphologicalComputationWS2 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS2(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) float64 {
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

	return continuous.KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, w1s1Indices, k, false) - continuous.KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, s1Indices, k, eta)
}

// MorphologicalComputationMI1 [...]
func MorphologicalComputationMI1(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger1(w2w1s1a1, w2Indices, w1Indices, k, false) - continuous.KraskovStoegbauerGrassberger1(w2w1s1a1, s1Indices, a1Indices, k, eta)
}

// MorphologicalComputationMI2 [...]
func MorphologicalComputationMI2(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger2(w2w1s1a1, w2Indices, w1Indices, k, false) - continuous.KraskovStoegbauerGrassberger2(w2w1s1a1, a1Indices, s1Indices, k, eta)
}

// MorphologicalComputationCA1 quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCA = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, eta) - continuous.KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationCA2 quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCA = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCA2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, eta) - continuous.KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
}
