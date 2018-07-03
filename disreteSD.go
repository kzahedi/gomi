package main

import (
	"fmt"
	"os"

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
	writeOutputSD(p, result, "MI_W")

}

func miaDiscreteSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Discrete SD")
	}
	w2a1w1 := makeW2A1W1Discrete(data, p)
	result := state.MorphologicalComputationA(w2a1w1)
	writeOutputSD(p, result, "MI_A")
}

func miaPrimeDiscreteSD(p Parameters, data Data) {

	if p.Verbose {
		fmt.Println("MI_A Prime Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	wbins := 1
	if len(p.WBins) > 0 {
		for _, v := range p.WBins {
			wbins *= v
		}
	} else {
		for i := 0; i < len(data.W[0]); i++ {
			wbins *= p.GlobalBins
		}
	}

	w2a1w1 := makeW2A1W1Discrete(data, p)
	result := state.MorphologicalComputationA(w2a1w1)

	// f := 1.0 / float64(len(result))
	// n := float64(wbins) / float64(len(result))

	for i, v := range result {
		result[i] = 1.0 - v
	}

	writeOutputSD(p, result, "MI_A_Prime")
}

func mimiDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_MI Prime Discrete SD")
	data.Discretise(p)
}

func misyDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Discrete SD")
	data.Discretise(p)
}

func micaDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_CA Prime Discrete SD")
	data.Discretise(p)
}

func miwaDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_WA Prime Discrete SD")
	data.Discretise(p)
}

func miwsDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_WS Prime Discrete SD")
	data.Discretise(p)
}

func miwpDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Discrete SD")
	data.Discretise(p)
}

func caDiscreteSD(p Parameters, data Data) {
	fmt.Println("CA Prime Discrete SD")
	data.Discretise(p)
}

func uiDiscreteSD(p Parameters, data Data) {
	fmt.Println("UI Prime Discrete SD")
	data.Discretise(p)
}

func ciDiscreteSD(p Parameters, data Data) {
	fmt.Println("CI Prime Discrete SD")
	data.Discretise(p)
}

func miinDiscreteSD(p Parameters, data Data) {
	fmt.Println("MI_IN Prime Discrete SD")
	data.Discretise(p)
}
