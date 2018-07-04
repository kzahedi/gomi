package main

import (
	"fmt"
	"math"

	"github.com/kzahedi/gomi/discrete/state"
)

func discreteSDCalculations(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwDiscreteSD(p, d)
	case "MI_A":
		miaDiscreteSD(p, d)
	case "MI_A_Prime":
		miaPrimeDiscreteSD(p, d)
	case "MI_MI":
		mimiDiscreteSD(p, d)
	case "MI_SY":
		misyDiscreteSD(p, d)
	case "MI_CA":
		micaDiscreteSD(p, d)
	case "MI_WA":
		miwaDiscreteSD(p, d)
	case "MI_WS":
		miwsDiscreteSD(p, d)
	case "MI_Wp":
		miwpDiscreteSD(p, d)
	case "CA":
		caDiscreteSD(p, d)
	case "UI":
		uiDiscreteSD(p, d)
	case "CI":
		ciDiscreteSD(p, d)
	case "MI_IN":
		miinDiscreteSD(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of discrete-state-dependent measures.", p.MeasureName))
	}
}

func miwDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_W Discrete SD")
	}

	w2w1a1 := makeW2W1A1Discrete(data, p)
	result := state.MorphologicalComputationW(w2w1a1)
	writeOutputSD(p, result, "MI_W discrete")

}

func miaDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Discrete SD")
	}
	w2a1w1 := makeW2A1W1Discrete(data, p)
	result := state.MorphologicalComputationA(w2a1w1)
	writeOutputSD(p, result, "MI_A discrete")
}

func miaPrimeDiscreteSD(p Parameters, data Data) {

	if p.Verbose {
		fmt.Println("MI_A Prime Discrete Avg")
	}

	wbins := calculateWBins(p, data)
	z := math.Log2(float64(wbins))

	w2a1w1 := makeW2A1W1Discrete(data, p)
	result := state.MorphologicalComputationA(w2a1w1)

	for i, v := range result {
		result[i] = 1.0 - v/z
	}

	writeOutputSD(p, result, "MI_A_Prime discrete")
}

func mimiDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_MI Prime Discrete SD")
	}

	w2w1 := makeW2W1Discrete(data, p)
	a1s1 := makeA1S1Discrete(data, p)

	result := state.MorphologicalComputationMI(w2w1, a1s1)
	writeOutputSD(p, result, "MI_MI discrete")
}

func micaDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_CA Discrete SD")
	}

	w2w1 := makeW2W1Discrete(data, p)
	w2a1 := makeW2A1Discrete(data, p)

	result := state.MorphologicalComputationCA(w2w1, w2a1)
	writeOutputSD(p, result, "MI_CA discrete")
}

func miwaDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Discrete SD")
	}

	w2w1a1 := makeW2W1A1Discrete(data, p)
	result := state.MorphologicalComputationWA(w2w1a1)
	writeOutputSD(p, result, "MI_WA discrete")

}

func miwsDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WS Discrete SD")
	}

	w2w1s1 := makeW2W1S1Discrete(data, p)
	result := state.MorphologicalComputationWS(w2w1s1)
	writeOutputSD(p, result, "MI_WS discrete")
}

func caDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("CA Prime Discrete SD")
	}
	w2w1 := makeW2W1Discrete(data, p)
	w2a1 := makeW2A1Discrete(data, p)

	result := state.MorphologicalComputationCA(w2w1, w2a1)
	writeOutputSD(p, result, "MI_CA discrete")
}

func miinDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_IN Prime Discrete SD")
	}
	abins := calculateABins(p, data)

	a1s1 := makeA1S1Discrete(data, p)
	result := state.MorphologicalComputationIN(a1s1, abins)
	writeOutputSD(p, result, "MI_IN discrete")
}

func uiDiscreteSD(p Parameters, data Data) {
	fmt.Println("UI Prime Discrete SD not implemented yet.")
}

func ciDiscreteSD(p Parameters, data Data) {
	fmt.Println("CI Prime Discrete SD not implemented yet.")
}
func misyDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Discrete SD is not implemented yet.")
}

func miwpDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Discrete SD")
}
