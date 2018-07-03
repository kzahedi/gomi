package main

import (
	"github.com/kzahedi/goent/dh"
	entropy "github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////////////////////////
// W2, W1, A1
////////////////////////////////////////////////////////////////////////////////

func makeW2W1A1Discrete(d Data, p Parameters) [][]int {

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

func makePW2W1A1(d Data, p Parameters) [][][]float64 {
	w2w1a1 := makeW2W1A1Discrete(d, p)
	pw2w1a1 := entropy.Emperical3D(w2w1a1)
	return pw2w1a1
}

func makeW2W1A1(d Data, p Parameters) ([][]float64, []int, []int, []int) {
	checkW(d)
	checkA(d)

	wDim := len(d.W)
	aDim := len(d.A)
	n := len(d.W[0]) - 1
	w2w1a1 := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		w2w1a1[i] = make([]float64, wDim+wDim+aDim, wDim+wDim+aDim)
		for wi := 0; wi < wDim; wi++ {
			w2w1a1[i][wi] = d.W[i+1][wi]
		}
		for wi := 0; wi < wDim; wi++ {
			w2w1a1[i][wDim+wi] = d.W[i][wi]
		}
		for ai := 0; ai < aDim; ai++ {
			w2w1a1[i][wDim+wDim+ai] = d.A[i][ai]
		}
	}

	var w2indices []int
	var w1indices []int
	var a1indices []int

	index := 0
	for wi := 0; wi < wDim; wi++ {
		w2indices = append(w2indices, index)
		index++
	}
	for wi := 0; wi < wDim; wi++ {
		w1indices = append(w1indices, index)
		index++
	}
	for ai := 0; ai < aDim; ai++ {
		a1indices = append(a1indices, index)
		index++
	}

	return w2w1a1, w2indices, w1indices, a1indices
}

////////////////////////////////////////////////////////////////////////////////
// W2, A1, W1
////////////////////////////////////////////////////////////////////////////////

func makeW2A1W1Discrete(d Data, p Parameters) [][]int {
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

func makePW2A1W1(d Data, p Parameters) [][][]float64 {
	w2a1w1 := makeW2A1W1Discrete(d, p)
	pw2a1w1 := entropy.Emperical3D(w2a1w1)
	return pw2a1w1
}

////////////////////////////////////////////////////////////////////////////////
// W2, W1
////////////////////////////////////////////////////////////////////////////////

func makeW2W1Discrete(d Data, p Parameters) [][]int {
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
		w2w1[i] = make([]int, 3, 3)
		w2w1[i][0] = w[i+1]
		w2w1[i][2] = w[i]
	}

	return w2w1
}

func makePW2W1(d Data, p Parameters) [][]float64 {
	w2w1 := makeW2W1Discrete(d, p)
	pw2w1 := entropy.Emperical2D(w2w1)
	return pw2w1
}

////////////////////////////////////////////////////////////////////////////////
// A1, S1
////////////////////////////////////////////////////////////////////////////////

func makeA1S1Discrete(d Data, p Parameters) [][]int {
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
		a1s1[i] = make([]int, 3, 3)
		a1s1[i][0] = a[i]
		a1s1[i][2] = s[i]
	}

	return a1s1
}

func makePA1S1(d Data, p Parameters) [][]float64 {
	a1s1 := makeA1S1Discrete(d, p)
	pa1s1 := entropy.Emperical2D(a1s1)
	return pa1s1
}

////////////////////////////////////////////////////////////////////////////////
// S2, S1, A1
////////////////////////////////////////////////////////////////////////////////

func makeS2S1A1Discrete(d Data, p Parameters) [][]int {
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

func makePS2S1A1(d Data, p Parameters) [][][]float64 {
	s2s1a1 := makeS2S1A1Discrete(d, p)
	ps2s1a1 := entropy.Emperical3D(s2s1a1)
	return ps2s1a1
}

////////////////////////////////////////////////////////////////////////////////
// W2, W1, S1
////////////////////////////////////////////////////////////////////////////////

func makeW2W1S1Discrete(d Data, p Parameters) [][]int {
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

func makePW2W1S1(d Data, p Parameters) [][][]float64 {
	w2w1s1 := makeW2W1S1Discrete(d, p)
	pw2w1s1 := entropy.Emperical3D(w2w1s1)
	return pw2w1s1
}

////////////////////////////////////////////////////////////////////////////////
// W2, A1
////////////////////////////////////////////////////////////////////////////////

func makeW2A1Discrete(d Data, p Parameters) [][]int {
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

func makePW2A1(d Data, p Parameters) [][]float64 {
	w2a1 := makeW2A1Discrete(d, p)
	pw2a1 := entropy.Emperical2D(w2a1)
	return pw2a1
}
