package gomi

import (
	"github.com/kzahedi/goent/dh"
	entropy "github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/sm"
)

////////////////////////////////////////////////////////////////////////////////
// W2, W1, A1
////////////////////////////////////////////////////////////////////////////////

// MakeW2W1A1Discrete returns a slice with (w',w,s,a).
// The retuned slice has four columns.
func MakeW2W1A1Discrete(d Data, p Parameters) [][]int {

	checkW(d)
	checkA(d)

	var wbins []int
	var abins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	if len(p.ABins) > 0 {
		abins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	w2w1a1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2w1a1[i] = make([]int, 3, 3)
		w2w1a1[i][0] = w[i+1]
		w2w1a1[i][1] = w[i]
		w2w1a1[i][2] = a[i]
	}

	return w2w1a1
}

// MakePW2W1A1 returns the joint distribution p(w',w,a)
func MakePW2W1A1(d Data, p Parameters) [][][]float64 {
	w2w1a1 := MakeW2W1A1Discrete(d, p)
	pw2w1a1 := entropy.Empirical3D(w2w1a1)
	return pw2w1a1
}

// MakePW2W1A1Sparse returns the joint distribution p(w',w,a) as SparseMatrix
func MakePW2W1A1Sparse(d Data, p Parameters) sm.SparseMatrix {
	w2w1a1 := MakeW2W1A1Discrete(d, p)
	pw2w1a1 := entropy.Empirical3DSparse(w2w1a1)
	return pw2w1a1
}

////////////////////////////////////////////////////////////////////////////////
// W2, A1, W1
////////////////////////////////////////////////////////////////////////////////

// MakeW2A1W1Discrete returns a slice with (w',a,w).
// The retuned slice has three columns.
func MakeW2A1W1Discrete(d Data, p Parameters) [][]int {
	checkW(d)
	checkA(d)

	var wbins []int
	var abins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	if len(p.ABins) > 0 {
		abins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	w2a1w1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2a1w1[i] = make([]int, 3, 3)
		w2a1w1[i][0] = w[i+1]
		w2a1w1[i][1] = a[i]
		w2a1w1[i][2] = w[i]
	}

	return w2a1w1
}

// MakePW2A1W1 return the joint distribution p(w',a,w)
func MakePW2A1W1(d Data, p Parameters) [][][]float64 {
	w2a1w1 := MakeW2A1W1Discrete(d, p)
	pw2a1w1 := entropy.Empirical3D(w2a1w1)
	return pw2a1w1
}

// MakePW2A1W1Sparse return the joint distribution p(w',a,w)
func MakePW2A1W1Sparse(d Data, p Parameters) sm.SparseMatrix {
	w2a1w1 := MakeW2A1W1Discrete(d, p)
	pw2a1w1 := entropy.Empirical3DSparse(w2a1w1)
	return pw2a1w1
}

////////////////////////////////////////////////////////////////////////////////
// W2, W1
////////////////////////////////////////////////////////////////////////////////

// MakeW2W1Discrete returns a slice with (w',w).
// The retuned slice has two columns.
func MakeW2W1Discrete(d Data, p Parameters) [][]int {
	checkW(d)

	var wbins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)

	w2w1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2w1[i] = make([]int, 2, 2)
		w2w1[i][0] = w[i+1]
		w2w1[i][1] = w[i]
	}

	return w2w1
}

// MakePW2W1 return the joint distribution p(w',w)
func MakePW2W1(d Data, p Parameters) [][]float64 {
	w2w1 := MakeW2W1Discrete(d, p)
	pw2w1 := entropy.Empirical2D(w2w1)
	return pw2w1
}

// MakePW2W1Sparse return the joint distribution p(w',w)
func MakePW2W1Sparse(d Data, p Parameters) sm.SparseMatrix {
	w2w1 := MakeW2W1Discrete(d, p)
	pw2w1 := entropy.Empirical2DSparse(w2w1)
	return pw2w1
}

////////////////////////////////////////////////////////////////////////////////
// A1, S1
////////////////////////////////////////////////////////////////////////////////

// MakeA1S1Discrete returns a slice with (a,s)
// The retuned slice has two columns.
func MakeA1S1Discrete(d Data, p Parameters) [][]int {
	checkA(d)
	checkS(d)

	var sbins []int
	var abins []int

	d.Discretise(p)

	if len(p.ABins) > 0 {
		abins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	if len(p.SBins) > 0 {
		sbins = p.SBins
	} else {
		for i := 0; i < len(d.Discretised.S[0]); i++ {
			sbins = append(sbins, p.GlobalBins)
		}
	}

	s := dh.MakeUnivariateRelabelled(d.Discretised.S, sbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	a1s1 := make([][]int, len(a), len(a))

	for i := 0; i < len(a); i++ {
		a1s1[i] = make([]int, 2, 2)
		a1s1[i][0] = a[i]
		a1s1[i][1] = s[i]
	}

	return a1s1
}

// MakePA1S1 return the joint distribution p(a,s)
func MakePA1S1(d Data, p Parameters) [][]float64 {
	a1s1 := MakeA1S1Discrete(d, p)
	pa1s1 := entropy.Empirical2D(a1s1)
	return pa1s1
}

// MakePA1S1Sparse return the joint distribution p(a,s)
func MakePA1S1Sparse(d Data, p Parameters) sm.SparseMatrix {
	a1s1 := MakeA1S1Discrete(d, p)
	pa1s1 := entropy.Empirical2DSparse(a1s1)
	return pa1s1
}

////////////////////////////////////////////////////////////////////////////////
// S2, S1, A1
////////////////////////////////////////////////////////////////////////////////

// MakeS2S1A1Discrete returns a slice with (s',s,a)
// The retuned slice has three columns.
func MakeS2S1A1Discrete(d Data, p Parameters) [][]int {
	checkS(d)
	checkA(d)

	var sbins []int
	var abins []int

	d.Discretise(p)

	if len(p.SBins) > 0 {
		sbins = p.SBins
	} else {
		for i := 0; i < len(d.Discretised.S[0]); i++ {
			sbins = append(sbins, p.GlobalBins)
		}
	}

	if len(p.ABins) > 0 {
		abins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	s := dh.MakeUnivariateRelabelled(d.Discretised.S, sbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	s2s1a1 := make([][]int, len(s)-1, len(s)-1)

	for i := 0; i < len(s)-1; i++ {
		s2s1a1[i] = make([]int, 3, 3)
		s2s1a1[i][0] = s[i+1]
		s2s1a1[i][1] = s[i]
		s2s1a1[i][2] = a[i]
	}

	return s2s1a1
}

// MakePS2S1A1 return the joint distribution p(s',s,a)
func MakePS2S1A1(d Data, p Parameters) [][][]float64 {
	s2s1a1 := MakeS2S1A1Discrete(d, p)
	ps2s1a1 := entropy.Empirical3D(s2s1a1)
	return ps2s1a1
}

////////////////////////////////////////////////////////////////////////////////
// W2, W1, S1
////////////////////////////////////////////////////////////////////////////////

// MakeW2W1S1Discrete returns a slice with (w',w,s)
// The retuned slice has three columns.
func MakeW2W1S1Discrete(d Data, p Parameters) [][]int {
	checkW(d)
	checkS(d)

	var wbins []int
	var sbins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	if len(p.SBins) > 0 {
		sbins = p.SBins
	} else {
		for i := 0; i < len(d.Discretised.S[0]); i++ {
			sbins = append(sbins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)
	s := dh.MakeUnivariateRelabelled(d.Discretised.S, sbins)

	w2w1s1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2w1s1[i] = make([]int, 3, 3)
		w2w1s1[i][0] = w[i+1]
		w2w1s1[i][1] = w[i]
		w2w1s1[i][2] = s[i]
	}

	return w2w1s1
}

// MakePW2W1S1 return the joint distribution p(w',w,s)
func MakePW2W1S1(d Data, p Parameters) [][][]float64 {
	w2w1s1 := MakeW2W1S1Discrete(d, p)
	pw2w1s1 := entropy.Empirical3D(w2w1s1)
	return pw2w1s1
}

////////////////////////////////////////////////////////////////////////////////
// W2, A1
////////////////////////////////////////////////////////////////////////////////

// MakeW2A1Discrete returns a slice with (w',a)
// The retuned slice has two columns.
func MakeW2A1Discrete(d Data, p Parameters) [][]int {
	checkW(d)
	checkA(d)

	var wbins []int
	var abins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	if len(p.ABins) > 0 {
		abins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	w2a1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2a1[i] = make([]int, 2, 2)
		w2a1[i][0] = w[i+1]
		w2a1[i][1] = a[i]
	}

	return w2a1
}

// MakePW2A1 return the joint distribution p(w',a)
func MakePW2A1(d Data, p Parameters) [][]float64 {
	w2a1 := MakeW2A1Discrete(d, p)
	pw2a1 := entropy.Empirical2D(w2a1)
	return pw2a1
}
