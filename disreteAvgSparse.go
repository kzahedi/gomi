package gomi

import (
	"fmt"
	"math"

	"github.com/kzahedi/gomi/discrete/sparse"
)

// DiscreteAvgCalculationsSparse ...
func DiscreteAvgCalculationsSparse(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwDiscreteAvgSparse(p, d)
	case "MI_A":
		miaDiscreteAvgSparse(p, d)
	case "MI_A_Prime":
		miaPrimeDiscreteAvgSparse(p, d)
	case "MI_MI":
		mimiDiscreteAvgSparse(p, d)
	case "MI_SY":
		// misyDiscreteAvgSparse(p, d)
	case "MI_SY_NID":
		// misynidDiscreteAvgSparse(p, d)
	case "MI_CA":
		// micaDiscreteAvgSparse(p, d)
	case "MI_WA":
		// miwaDiscreteAvgSparse(p, d)
	case "MI_WS":
		// miwsDiscreteAvgSparse(p, d)
	case "MI_Wp":
		// miwpDiscreteAvgSparse(p, d)
	case "CA":
		// caDiscreteAvgSparse(p, d)
	case "UI":
		// uiDiscreteAvgSparse(p, d)
	case "CI":
		// ciDiscreteAvgSparse(p, d)
	case "MI_IN":
		// miinDiscreteAvgSparse(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of discrete-avg measures.", p.MeasureName))
	}
}

func miwDiscreteAvgSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_W Discrete Avg")
	}

	pw2w1a1 := MakePW2W1A1Sparse(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationW(pw2w1a1)

	writeOutputAvg(p, result, "MI_W discrete (sparse matrix)", output)
}

func miaDiscreteAvgSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_A Discrete Avg")
	}

	pw2a1w1 := MakePW2A1W1Sparse(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationA(pw2a1w1)

	writeOutputAvg(p, result, "MI_A discrete (sparse matrix)", output)
}

func miaPrimeDiscreteAvgSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_A Prime Discrete Avg")
	}

	wBins := CalculateWBins(p, data)

	pw2a1w1 := MakePW2A1W1Sparse(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := 1.0 - sparse.MorphologicalComputationA(pw2a1w1)/math.Log2(float64(wBins))

	writeOutputAvg(p, result, "MI_A_Prime discrete (sparse matrix)", output)
}

func mimiDiscreteAvgSparse(p Parameters, data Data) {
	var output Output
	if p.Verbose {
		fmt.Println("MI_MI Discrete Avg")
	}

	pw2w1 := MakePW2W1Sparse(data, p)
	pa1s1 := MakePA1S1Sparse(data, p)

	if p.Verbose == true {
		fmt.Println(p)
	}

	result := sparse.MorphologicalComputationMI(pw2w1, pa1s1)
	writeOutputAvg(p, result, "MI_MI discrete (sparse matrix)", output)
	// TODO: results look wrong
}

// func misyDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_SY Discrete Avg")
// 	}

// 	pw2a1w1 := MakePW2A1W1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationSY(pw2a1w1, p.Iterations, p.Verbose)

// 	writeOutputAvg(p, result, "MI_SY discrete", output)
// }

// func misynidDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_SY_NID Discrete Avg")
// 	}

// 	pw2a1w1 := MakePW2A1W1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationSyNid(pw2a1w1, p.Iterations, p.Verbose)

// 	writeOutputAvg(p, result, "MI_SY_NID discrete", output)
// }

// func miwaDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_WA Discrete Avg")
// 	}

// 	pw2w1a1 := MakePW2W1A1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationWA(pw2w1a1)

// 	writeOutputAvg(p, result, "MI_WA discrete", output)
// }

// func miwsDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_WS Discrete Avg")
// 	}

// 	pw2w1s1 := MakePW2W1S1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationWS(pw2w1s1)

// 	writeOutputAvg(p, result, "MI_WS discrete", output)
// }

// func miwpDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_Wp Discrete Avg")
// 	}

// 	pw2a1w1 := MakePW2A1W1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationWp(pw2a1w1, p.Iterations, p.Verbose)

// 	writeOutputAvg(p, result, "MI_Wp discrete", output)
// }

// func caDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("CA Discrete Avg")
// 	}

// 	ps2s1a1 := MakePS2S1A1(data, p)

// 	sBins := CalculateSBins(p, data)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationIntrinsicCA(ps2s1a1, sBins)
// 	writeOutputAvg(p, result, "CA discrete", output)
// }

// func miinDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_IN Discrete Avg")
// 	}

// 	aBins := CalculateABins(p, data)

// 	pa1s1 := MakePA1S1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationIN(pa1s1, aBins)
// 	writeOutputAvg(p, result, "MI_IN discrete", output)
// }

// func micaDiscreteAvgSparse(p Parameters, data Data) {
// 	var output Output
// 	if p.Verbose {
// 		fmt.Println("MI_CA Discrete Avg")
// 	}

// 	pw2w1 := MakePW2W1(data, p)
// 	pw2a1 := MakePW2A1(data, p)

// 	if p.Verbose == true {
// 		fmt.Println(p)
// 	}

// 	result := discrete.MorphologicalComputationCA(pw2w1, pw2a1)
// 	writeOutputAvg(p, result, "MI_CA discrete", output)
// }

// func uiDiscreteAvgSparse(p Parameters, data Data) {
// 	fmt.Println("UI Discrete Avg is not implemented for discrete data yet.")
// }

// func ciDiscreteAvgSparse(p Parameters, data Data) {
// 	fmt.Println("CI Discrete Avg is not implemented for discrete data yet.")
// }
