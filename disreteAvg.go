package main

import (
	"fmt"
	"math"
	"os"

	"github.com/kzahedi/gomi/discrete"
)

func discreteAvgCalculations(p Parameters, d Data) {
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

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2w1a1 := makePW2W1A1(data, p)
	result := discrete.MorphologicalComputationW(pw2w1a1)

	writeOutputAvg(p, result, "MI_W")
}

func miaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2a1w1 := makePW2A1W1(data, p)
	result := discrete.MorphologicalComputationA(pw2a1w1)

	writeOutputAvg(p, result, "MI_A")
}

func miaPrimeDiscreteAvg(p Parameters, data Data) {
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

	pw2a1w1 := makePW2A1W1(data, p)
	result := 1.0 - discrete.MorphologicalComputationA(pw2a1w1)/math.Log2(float64(wbins))

	writeOutputAvg(p, result, "MI_A_Prime")
}

func mimiDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_MI Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.S) == 0 {
		fmt.Print("S is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2w1 := makePW2W1(data, p)
	pa1s1 := makePA1S1(data, p)

	result := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
	writeOutputAvg(p, result, "MI_MI")
	// TODO: results look wrong
}

func misyDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_SY Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2a1w1 := makePW2A1W1(data, p)
	result := discrete.MorphologicalComputationSY(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_SY")
}

func misynidDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_SY_NID Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2a1w1 := makePW2A1W1(data, p)
	fmt.Println(fmt.Sprintf("%t", p.Verbose))
	result := discrete.MorphologicalComputationSyNid(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_SY_NID")
}

func miwaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2w1a1 := makePW2W1A1(data, p)
	result := discrete.MorphologicalComputationWA(pw2w1a1)

	writeOutputAvg(p, result, "MI_WA")
}

func miwsDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WS Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.S) == 0 {
		fmt.Print("S is empty")
		os.Exit(0)
	}

	pw2w1s1 := makePW2W1S1(data, p)
	result := discrete.MorphologicalComputationWS(pw2w1s1)

	writeOutputAvg(p, result, "MI_WS")
}

func miwpDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_Wp Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2a1w1 := makePW2A1W1(data, p)
	result := discrete.MorphologicalComputationWp(pw2a1w1, p.Iterations, p.Verbose)

	writeOutputAvg(p, result, "MI_Wp")

}

func caDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("CA Discrete Avg")
	}

	if len(data.S) == 0 {
		fmt.Print("S is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	ps2s1a1 := makePS2S1A1(data, p)

	sbins := p.GlobalBins
	if len(p.SBins) > 0 {
		sbins = 1
		for _, v := range p.SBins {
			sbins = sbins * v
		}
	}

	result := discrete.MorphologicalComputationIntrinsicCA(ps2s1a1, sbins)
	writeOutputAvg(p, result, "CA")
}

func uiDiscreteAvg(p Parameters, data Data) {
	fmt.Println("UI Discrete Avg is not implemented for discrete data yet.")
}

func ciDiscreteAvg(p Parameters, data Data) {
	fmt.Println("CI Discrete Avg is not implemented for discrete data yet.")
}

func miinDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_IN Discrete Avg")
	}

	if len(data.S) == 0 {
		fmt.Print("S is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pa1s1 := makePA1S1(data, p)

	abins := p.GlobalBins
	if len(p.ABins) > 0 {
		abins = 1
		for _, v := range p.ABins {
			abins = abins * v
		}
	}

	result := discrete.MorphologicalComputationIN(pa1s1, abins)
	writeOutputAvg(p, result, "MI_IN")

}

func micaDiscreteAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_CA Discrete Avg")
	}

	if len(data.W) == 0 {
		fmt.Print("W is empty")
		os.Exit(0)
	}

	if len(data.A) == 0 {
		fmt.Print("A is empty")
		os.Exit(0)
	}

	pw2w1 := makePW2W1(data, p)
	pw2a1 := makePW2A1(data, p)

	result := discrete.MorphologicalComputationCA(pw2w1, pw2a1)
	writeOutputAvg(p, result, "MI_CA")
}
