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

	writeOutputAvg(p, result, "MI_A")

}

func mimiDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_MI Prime Discrete Avg")

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
	fmt.Println("MI_SY Prime Discrete Avg")
	data.Discretise(p)
}

func misynidDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_SY_NID Prime Discrete Avg")
	data.Discretise(p)
}

func micaDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_CA Prime Discrete Avg")
	data.Discretise(p)
}

func miwaDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_WA Prime Discrete Avg")
	data.Discretise(p)
}

func miwsDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_WS Prime Discrete Avg")
	data.Discretise(p)
}

func miwpDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Discrete Avg")
	data.Discretise(p)
}

func caDiscreteAvg(p Parameters, data Data) {
	fmt.Println("CA Prime Discrete Avg")
	data.Discretise(p)
}

func uiDiscreteAvg(p Parameters, data Data) {
	fmt.Println("UI Prime Discrete Avg")
	data.Discretise(p)
}

func ciDiscreteAvg(p Parameters, data Data) {
	fmt.Println("CI Prime Discrete Avg")
	data.Discretise(p)
}

func miinDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_IN Prime Discrete Avg")
	data.Discretise(p)
}
