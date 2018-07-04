package state

import (
	"math"

	"github.com/kzahedi/goent/discrete/state"
)

// MorphologicalComputationW quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
// and
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
//   MC_W = I(W';W|A)
func MorphologicalComputationW(w2w1a1 [][]int) []float64 {
	return state.ConditionalMutualInformationBase2(w2w1a1)
}

// MorphologicalComputationA quantifies morphological computation as the information that is contained in
// A about W' that is not contained in W. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
//   MC_W = I(W';A|W)
func MorphologicalComputationA(w2a1w1 [][]int) []float64 {
	return state.ConditionalMutualInformationBase2(w2a1w1)
}

// MorphologicalComputationCW quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
//   MC_CW = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCW(w2w1a1 [][]int) []float64 {
	w2w1 := make([][]int, len(w2w1a1), len(w2w1a1))
	w2a1 := make([][]int, len(w2w1a1), len(w2w1a1))
	for i := 0; i < len(w2w1a1); i++ {
		w2w1[i] = make([]int, 2, 2)
		w2a1[i] = make([]int, 2, 2)
		w2w1[i][0] = w2w1a1[i][0]
		w2w1[i][1] = w2w1a1[i][1]
		w2a1[i][0] = w2w1a1[i][0]
		w2a1[i][1] = w2w1a1[i][2]
	}
	r1 := state.MutualInformationBase2(w2w1)
	r2 := state.MutualInformationBase2(w2a1)
	r := make([]float64, len(r1), len(r1))
	for i := 0; i < len(r1); i++ {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationWA = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA(w2w1a1 [][]int) []float64 {
	w2a1 := make([][]int, len(w2w1a1), len(w2w1a1))
	for i := 0; i < len(w2w1a1); i++ {
		w2a1[i] = make([]int, 2, 2)
		w2a1[i][0] = w2w1a1[i][0]
		w2a1[i][1] = w2w1a1[i][2]
	}

	r1 := state.ConditionalMutualInformationBase2(w2w1a1)
	r2 := state.MutualInformationBase2(w2a1)
	r := make([]float64, len(w2w1a1), len(w2w1a1))

	for i := 0; i < len(r1); i++ {
		r[i] = r1[i] - r2[i]
	}

	return r
}

// MorphologicalComputationWS = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS(w2w1s1 [][]int) []float64 {
	w2s1 := make([][]int, len(w2w1s1), len(w2w1s1))
	for i := 0; i < len(w2w1s1); i++ {
		w2s1[i] = make([]int, 2, 2)
		w2s1[i][0] = w2w1s1[i][0]
		w2s1[i][1] = w2w1s1[i][2]
	}

	r1 := state.ConditionalMutualInformationBase2(w2w1s1)
	r2 := state.MutualInformationBase2(w2s1)
	r := make([]float64, len(w2w1s1), len(w2w1s1))

	for i := 0; i < len(r1); i++ {
		r[i] = r1[i] - r2[i]
	}

	return r
}

// MorphologicalComputationMI quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
//   MC_MI = I(W';W) - I(A;S)
func MorphologicalComputationMI(w2w1 [][]int, a1s1 [][]int) []float64 {
	r1 := state.MutualInformationBase2(w2w1)
	r2 := state.MutualInformationBase2(a1s1)
	// fmt.Println(r1)
	// fmt.Println(r2)
	r := make([]float64, len(r1), len(r1))
	for i := 0; i < len(r1); i++ {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationCA quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCW = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCA(w2w1, w2a1 [][]int) []float64 {
	r1 := state.MutualInformationBase2(w2w1)
	r2 := state.MutualInformationBase2(w2a1)
	n := len(r1)
	r := make([]float64, n, n)
	for i := 0; i < n; i++ {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationIN quantifies morphological computation as the in-sourcable
// complexity of the world process.
func MorphologicalComputationIN(a1s1 [][]int, abins int) []float64 {
	r := state.MutualInformationBase2(a1s1)
	l := math.Log2(float64(abins))
	for i, v := range r {
		r[i] = l - v
	}
	return r
}
