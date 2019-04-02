package discrete

import (
	"math"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/sm"
	stat "gonum.org/v1/gonum/stat"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// MorphologicalComputationWSparse quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
// and
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
func MorphologicalComputationWSparse(pw2w1a1 sm.SparseMatrix) float64 {
	return discrete.ConditionalMutualInformationBase2Sparse(pw2w1a1)
}

// MorphologicalComputationASparse quantifies morphological computation as the information that is contained in
// A about W' that is not contained in W. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationASparse(pw2a1w1 sm.SparseMatrix) float64 {
	return discrete.ConditionalMutualInformationBase2Sparse(pw2a1w1)
}

// MorphologicalComputationCASparse quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCA = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCASparse(pw2w1, pw2a1 sm.SparseMatrix) float64 {
	return discrete.MutualInformationBase2Sparse(pw2w1) - discrete.MutualInformationBase2Sparse(pw2a1)
}

// MorphologicalComputationWASparse = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWASparse(pw2w1a1 sm.SparseMatrix) float64 {
	pw2a1 := sm.CreateSparseMatrix()
	for _, index := range pw2w1a1.Indices {
		v, _ := pw2w1a1.Get(index)
		w2a1 := sm.SparseMatrixIndex{index[0], index[2]}
		pw2a1.Add(w2a1, v)
	}

	return discrete.ConditionalMutualInformationBase2Sparse(pw2w1a1) - discrete.MutualInformationBase2Sparse(pw2a1)
}

// MorphologicalComputationWSSparse = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWSSparse(pw2w1s1 sm.SparseMatrix) float64 {
	pw2s1 := sm.CreateSparseMatrix()
	for _, index := range pw2w1s1.Indices {
		v, _ := pw2w1s1.Get(index)
		w2s1 := sm.SparseMatrixIndex{index[0], index[2]}
		pw2s1.Add(w2s1, v)
	}

	return discrete.ConditionalMutualInformationBase2Sparse(pw2w1s1) - discrete.MutualInformationBase2Sparse(pw2s1)
}

// MorphologicalComputationMISparse quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
func MorphologicalComputationMISparse(pw2w1 sm.SparseMatrix, pa1s1 sm.SparseMatrix) float64 {
	return discrete.MutualInformationBase2Sparse(pw2w1) - discrete.MutualInformationBase2Sparse(pa1s1)
}

// MorphologicalComputationSYSparse quantifies morphological computation as the synergistic information that
// W and A contain about W'. For more details, please read
// TODO Paper reference
func MorphologicalComputationSYSparse(pw2w1a1 [][][]float64, iterations int, eta bool) float64 {
	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	split.NrOfStates = []int{w2Dim, w1Dim, a1Dim}

	split.CreateAlphabet()

	split.PTarget = make([]float64, w2Dim*w1Dim*a1Dim, w2Dim*w1Dim*a1Dim)
	for i, a := range split.Alphabet {
		split.PTarget[i] = pw2w1a1[a[0]][a[1]][a[2]]
	}

	split.Features = make(map[string][]int)
	split.Features["W,W'"] = []int{0, 2}
	split.Features["A,W'"] = []int{1, 2}
	split.Features["W,A"] = []int{0, 1}

	split.Init()
	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(iterations)
	}

	for i := 0; i < iterations; i++ {
		split.Iterate()
		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Finish()
	}

	return stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
}

// MorphologicalComputationSyNidSparse quantifies morphological computation as the synergistic
// information that W and A contain about W', excluding the input distribution
// (W,A). For more details, please read
// TODO Paper reference
func MorphologicalComputationSyNidSparse(pw2w1a1 [][][]float64, iterations int, eta bool) float64 {
	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	split.NrOfStates = []int{w2Dim, w1Dim, a1Dim}

	split.CreateAlphabet()

	split.PTarget = make([]float64, w2Dim*w1Dim*a1Dim, w2Dim*w1Dim*a1Dim)
	for i, a := range split.Alphabet {
		split.PTarget[i] = pw2w1a1[a[0]][a[1]][a[2]]
	}

	split.Features = make(map[string][]int)
	split.Features["W,W'"] = []int{0, 2}
	split.Features["A,W'"] = []int{1, 2}

	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(iterations)
	}

	split.Init()
	for i := 0; i < iterations; i++ {
		split.Iterate()
		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Finish()
	}

	return stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
}

// MorphologicalComputationWpSparse calculates the unique information W -> W'. For more details,
// please see
// Ghazi-Zahedi, Keyan and Langer, Carlotta and Ay, Nihat,
// Morphological Computation: Synergy of Body and Brain, Entropy, 2017
// func MorphologicalComputationWpSparse(pw2w1a1 sm.SparseMatrix, iterations int, eta bool) float64 {
// 	return MorphologicalComputationWSparse(pw2w1a1) - MorphologicalComputationSYSparse(pw2w1a1, iterations, eta)
// }

// MorphologicalComputationIntrinsicCASparse [...]
// For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationIntrinsicCASparse(ps2s1a1 [][][]float64, sbins int) float64 {
	s2Dim := len(ps2s1a1)
	s1Dim := len(ps2s1a1[0])
	a1Dim := len(ps2s1a1[0][0])
	ps1a1 := discrete.Create2D(s1Dim, a1Dim)
	ps2doa1 := discrete.Create2D(s2Dim, a1Dim)
	ps2dos1 := discrete.Create2D(s2Dim, s1Dim)
	pa1Cs1 := discrete.Create2D(a1Dim, s1Dim)
	ps1 := make([]float64, s1Dim, s1Dim)

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps1a1[s1][a1] += ps2s1a1[s2][s1][a1]
			}
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			ps1[s1] += ps1a1[s1][a1]
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			if ps1[s1] != 0.0 {
				if ps1[s1] > 0.0 {
					pa1Cs1[a1][s1] = ps1a1[s1][a1] / ps1[s1]
				}
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			for s1 := 0; s1 < s1Dim; s1++ {
				if ps1a1[s1][a1] > 0.0 {
					ps2doa1[s2][a1] += ps1[s1] * ps2s1a1[s2][s1][a1] / ps1a1[s1][a1]
				}
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2dos1[s2][s1] += pa1Cs1[a1][s1] * ps2doa1[s2][a1]
			}
		}
	}

	r := 0.0

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				if ps2dos1[s2][s1] > 0.0 && ps2doa1[s2][a1] > 0.0 {
					r += ps1a1[s1][a1] * ps2doa1[s2][a1] * math.Log2(ps2doa1[s2][a1]/ps2dos1[s2][s1])
				}
			}
		}
	}

	return 1.0 - r/math.Log2(float64(sbins))
}

// MorphologicalComputationIntrinsicCW [...]
// For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationIntrinsicCWSparse(ps2s1a1 [][][]float64) (r float64) {
	s2Dim := len(ps2s1a1)
	s1Dim := len(ps2s1a1[0])
	a1Dim := len(ps2s1a1[0][0])
	ps2Cs1 := discrete.Create2D(s2Dim, s1Dim)
	ps2Cs1Hat := discrete.Create2D(s2Dim, s1Dim)
	ps2Cs1a1 := discrete.Create3D(s2Dim, s1Dim, a1Dim)
	pa1Cs1 := discrete.Create2D(s2Dim, a1Dim)
	ps1a1 := discrete.Create2D(s1Dim, a1Dim)
	ps1 := make([]float64, s1Dim, s1Dim)
	pa1 := make([]float64, a1Dim, a1Dim)

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps1a1[s1][a1] += ps2s1a1[s2][s1][a1]
			}
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			ps1[s1] += ps1a1[s1][a1]
			pa1[a1] += ps1a1[s1][a1]
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			pa1Cs1[a1][s1] = ps1a1[s1][a1] / ps1[s1]
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2Cs1a1[s2][s1][a1] = ps2s1a1[s2][s1][a1] / ps1a1[s1][a1]
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2Cs1[s2][s1] += ps2Cs1a1[s2][s1][a1] * pa1Cs1[a1][s1]
				for s3 := 0; s3 < s1Dim; s3++ {
					ps2Cs1Hat[s2][s1] += pa1Cs1[a1][s1] * ps2Cs1a1[s2][s3][a1] * pa1Cs1[a1][s3] * ps1[s3] / pa1[a1]
				}
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			r += ps2Cs1[s2][s1] * math.Log2(ps2Cs1[s2][s1]/ps2Cs1Hat[s2][s1])
		}
	}

	return r
}

// MorphologicalComputationINSparse quantifies morphological computation as the in-sourcable
// complexity of the world process.
// func MorphologicalComputationINSparse(pa1s1 [][]float64, abins int) float64 {
// 	return math.Log2(float64(abins)) - discrete.MutualInformationBase2Sparse(pa1s1)
// }
