package gomi

import (
	"fmt"
	"math"

	"github.com/kzahedi/gomi/discrete/state/sparse"
)

// DiscreteSDCalculationsSparse ...
func DiscreteSDCalculationsSparse(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwDiscreteSDSparse(p, d)
	case "MI_A":
		miaDiscreteSDSparse(p, d)
	case "MI_A_Prime":
		miaPrimeDiscreteSDSparse(p, d)
	case "MI_MI":
		mimiDiscreteSDSparse(p, d)
	case "MI_SY":
		misyDiscreteSDSparse(p, d)
	case "MI_CA":
		micaDiscreteSDSparse(p, d)
	case "MI_WA":
		miwaDiscreteSDSparse(p, d)
	case "MI_WS":
		miwsDiscreteSDSparse(p, d)
	case "MI_Wp":
		miwpDiscreteSDSparse(p, d)
	case "CA":
		caDiscreteSDSparse(p, d)
	case "UI":
		uiDiscreteSDSparse(p, d)
	case "CI":
		ciDiscreteSDSparse(p, d)
	case "MI_IN":
		miinDiscreteSDSparse(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of discrete-state-dependent measures.", p.MeasureName))
	}
}

func miwDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_W Discrete SD")
	}

	w2w1a1 := MakeW2W1A1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationW(w2w1a1)
	writeOutputSD(p, result, "MI_W discrete (sparse matrix)", output)

}

func miaDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_A Discrete SD")
	}

	w2a1w1 := MakeW2A1W1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationA(w2a1w1)
	writeOutputSD(p, result, "MI_A discrete (sparse matrix)", output)
}

func miaPrimeDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_A Prime Discrete Avg")
	}

	wBins := CalculateWBins(p, data)
	z := math.Log2(float64(wBins))

	w2a1w1 := MakeW2A1W1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationA(w2a1w1)

	for i, v := range result {
		result[i] = 1.0 - v/z
	}

	writeOutputSD(p, result, "MI_A_Prime discrete (sparse matrix)", output)
}

func mimiDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_MI Prime Discrete SD")
	}

	w2w1 := MakeW2W1Discrete(data, p)
	a1s1 := MakeA1S1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationMI(w2w1, a1s1)
	writeOutputSD(p, result, "MI_MI discrete (sparse matrix)", output)
}

func micaDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_CA Discrete SD")
	}

	w2w1 := MakeW2W1Discrete(data, p)
	w2a1 := MakeW2A1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationCA(w2w1, w2a1)
	writeOutputSD(p, result, "MI_CA discrete (sparse matrix)", output)
}

func miwaDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_WA Discrete SD")
	}

	w2w1a1 := MakeW2W1A1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationWA(w2w1a1)
	writeOutputSD(p, result, "MI_WA discrete (sparse matrix)", output)

}

func miwsDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_WS Discrete SD")
	}

	w2w1s1 := MakeW2W1S1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationWS(w2w1s1)
	writeOutputSD(p, result, "MI_WS discrete (sparse matrix)", output)
}

func caDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("CA Prime Discrete SD")
	}

	w2w1 := MakeW2W1Discrete(data, p)
	w2a1 := MakeW2A1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationCA(w2w1, w2a1)
	writeOutputSD(p, result, "MI_CA discrete (sparse matrix)", output)
}

func miinDiscreteSDSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_IN Prime Discrete SD")
	}

	aBins := CalculateABins(p, data)
	a1s1 := MakeA1S1Discrete(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationIN(a1s1, aBins)
	writeOutputSD(p, result, "MI_IN discrete (sparse matrix)", output)
}

func uiDiscreteSDSparse(p Parameters, data Data) {
	fmt.Println("UI Prime Discrete SD not implemented yet.")
}

func ciDiscreteSDSparse(p Parameters, data Data) {
	fmt.Println("CI Prime Discrete SD not implemented yet.")
}
func misyDiscreteSDSparse(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Discrete SD is not implemented yet.")
}

func miwpDiscreteSDSparse(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Discrete SD")
}
