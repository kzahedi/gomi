package gomi

import (
	"fmt"
	"math"

	"github.com/kzahedi/gomi/discrete"
)

func DiscreteAvgCalculations(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwDiscreteAvg(p, d)
	case "MI_A":
		miaDiscreteAvg(p, d)
	case "MI_A_Prime":
		miaPrimeDiscreteAvg(p, d)
	case "MI_MI":
		mimiDiscreteAvg(p, d)
	case "MI_SY":
		misyDiscreteAvg(p, d)
	case "MI_SY_NID":
		misynidDiscreteAvg(p, d)
	case "MI_CA":
		micaDiscreteAvg(p, d)
	case "MI_WA":
		miwaDiscreteAvg(p, d)
	case "MI_WS":
		miwsDiscreteAvg(p, d)
	case "MI_Wp":
		miwpDiscreteAvg(p, d)
	case "CA":
		caDiscreteAvg(p, d)
	case "UI":
		uiDiscreteAvg(p, d)
	case "CI":
		ciDiscreteAvg(p, d)
	case "MI_IN":
		miinDiscreteAvg(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of discrete-avg measures.", p.MeasureName))
	}
}

func miwDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_W Discrete Avg")
	}

	pw2w1a1 := MakePW2W1A1(data, p)
	result := discrete.MorphologicalComputationW(pw2w1a1)

	writeOutputAvg(p, result, "MI_W discrete")
}

func miaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Discrete Avg")
	}

	pw2a1w1 := MakePW2A1W1(data, p)
	result := discrete.MorphologicalComputationA(pw2a1w1)

	writeOutputAvg(p, result, "MI_A discrete")
}

func miaPrimeDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Prime Discrete Avg")
	}

	wbins := CalculateWBins(p, data)

	pw2a1w1 := MakePW2A1W1(data, p)
	result := 1.0 - discrete.MorphologicalComputationA(pw2a1w1)/math.Log2(float64(wbins))

	writeOutputAvg(p, result, "MI_A_Prime discrete")
}

func mimiDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_MI Discrete Avg")
	}

	pw2w1 := MakePW2W1(data, p)
	pa1s1 := MakePA1S1(data, p)

	result := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
	writeOutputAvg(p, result, "MI_MI discrete")
	// TODO: results look wrong
}

func misyDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_SY Discrete Avg")
	}

	pw2a1w1 := MakePW2A1W1(data, p)
	result := discrete.MorphologicalComputationSY(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_SY discrete")
}

func misynidDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_SY_NID Discrete Avg")
	}

	pw2a1w1 := MakePW2A1W1(data, p)
	result := discrete.MorphologicalComputationSyNid(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_SY_NID discrete")
}

func miwaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Discrete Avg")
	}

	pw2w1a1 := MakePW2W1A1(data, p)
	result := discrete.MorphologicalComputationWA(pw2w1a1)

	writeOutputAvg(p, result, "MI_WA discrete")
}

func miwsDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WS Discrete Avg")
	}

	pw2w1s1 := MakePW2W1S1(data, p)
	result := discrete.MorphologicalComputationWS(pw2w1s1)

	writeOutputAvg(p, result, "MI_WS discrete")
}

func miwpDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_Wp Discrete Avg")
	}

	pw2a1w1 := MakePW2A1W1(data, p)
	result := discrete.MorphologicalComputationWp(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_Wp discrete")
}

func caDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("CA Discrete Avg")
	}

	ps2s1a1 := MakePS2S1A1(data, p)

	sBins := CalculateSBins(p, data)

	result := discrete.MorphologicalComputationIntrinsicCA(ps2s1a1, sBins)
	writeOutputAvg(p, result, "CA discrete")
}

func miinDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_IN Discrete Avg")
	}

	aBins := CalculateABins(p, data)

	pa1s1 := MakePA1S1(data, p)
	result := discrete.MorphologicalComputationIN(pa1s1, aBins)
	writeOutputAvg(p, result, "MI_IN discrete")
}

func micaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_CA Discrete Avg")
	}

	pw2w1 := MakePW2W1(data, p)
	pw2a1 := MakePW2A1(data, p)

	result := discrete.MorphologicalComputationCA(pw2w1, pw2a1)
	writeOutputAvg(p, result, "MI_CA discrete")
}

func uiDiscreteAvg(p Parameters, data Data) {
	fmt.Println("UI Discrete Avg is not implemented for discrete data yet.")
}

func ciDiscreteAvg(p Parameters, data Data) {
	fmt.Println("CI Discrete Avg is not implemented for discrete data yet.")
}
