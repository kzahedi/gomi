package main

////////////////////////////////////////////////////////////////////////////////
// W2, W1, S1, A1 continuous
////////////////////////////////////////////////////////////////////////////////

func MakeW2W1S1A1(d Data, p Parameters) ([][]float64, []int, []int, []int, []int) {
	checkW(d)
	checkA(d)
	checkS(d)

	wDim := len(d.W[0])
	aDim := len(d.A[0])
	sDim := len(d.S[0])
	n := len(d.W) - 1
	m := 2*wDim + sDim + aDim
	w2w1s1a1 := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		w2w1s1a1[i] = make([]float64, m, m)
		for wi := 0; wi < wDim; wi++ {
			w2w1s1a1[i][wi] = d.W[i+1][wi]
		}
		for wi := 0; wi < wDim; wi++ {
			w2w1s1a1[i][wDim+wi] = d.W[i][wi]
		}
		for si := 0; si < sDim; si++ {
			w2w1s1a1[i][wDim+wDim+si] = d.S[i][si]
		}
		for ai := 0; ai < aDim; ai++ {
			w2w1s1a1[i][wDim+wDim+sDim+ai] = d.A[i][ai]
		}
	}

	var w2indices []int
	var w1indices []int
	var s1indices []int
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
	for si := 0; si < sDim; si++ {
		s1indices = append(s1indices, index)
		index++
	}
	for ai := 0; ai < aDim; ai++ {
		a1indices = append(a1indices, index)
		index++
	}

	return w2w1s1a1, w2indices, w1indices, s1indices, a1indices
}

////////////////////////////////////////////////////////////////////////////////
// W2, W1, A1 continuous
////////////////////////////////////////////////////////////////////////////////

func MakeW2W1A1(d Data, p Parameters) ([][]float64, []int, []int, []int) {
	checkW(d)
	checkA(d)

	wDim := len(d.W[0])
	aDim := len(d.A[0])
	n := len(d.W) - 1
	m := 2*wDim + aDim
	w2w1a1 := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		w2w1a1[i] = make([]float64, m, m)
		for wi := 0; wi < wDim; wi++ {
			w2w1a1[i][wi] = d.W[i+1][wi]
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
