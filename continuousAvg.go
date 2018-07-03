package main

import (
	"fmt"

	"github.com/kzahedi/gomi/continuous"
)

func continuousAvgCalculations(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwContinuousAvg(p, d)
	case "MI_A":
		miaContinuousAvg(p, d)
	case "MI_A_Prime":
		miaPrimeContinuousAvg(p, d)
	case "MI_MI":
		mimiContinuousAvg(p, d)
	case "MI_SY":
		misyContinuousAvg(p, d)
	case "MI_CA":
		micaContinuousAvg(p, d)
	case "MI_WA":
		miwaContinuousAvg(p, d)
	case "MI_WS":
		miwsContinuousAvg(p, d)
	case "MI_Wp":
		miwpContinuousAvg(p, d)
	case "CA":
		caContinuousAvg(p, d)
	case "UI":
		uiContinuousAvg(p, d)
	case "CI":
		ciContinuousAvg(p, d)
	case "MI_IN":
		miinContinuousAvg(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of continuous-avg measures.", p.MeasureName))
	}
}

func miwContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_W Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := makeW2W1A1(data, p)
	result := continuous.MorphologicalComputationW(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_W continuous")

}

func miaContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := makeW2W1A1(data, p)
	result := continuous.MorphologicalComputationA(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_W continuous")
}

func mimiContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_MI Prime Continuous Avg")
	}

	w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices := makeW2W1S1A1(data, p)

	switch p.ContinuousMode {
	case 1:
		result := continuous.MorphologicalComputationMI1(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_MI continuous (KSG 1 Estimator)")
	case 2:
		result := continuous.MorphologicalComputationMI2(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_MI continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func micaContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_CA Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := makeW2W1A1(data, p)

	switch p.ContinuousMode {
	case 1:
		result := continuous.MorphologicalComputationCA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_CA continuous (KSG 1 Estimator)")
	case 2:
		result := continuous.MorphologicalComputationCA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_CA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func miwaContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := makeW2W1A1(data, p)

	switch p.ContinuousMode {
	case 1:
		result := continuous.MorphologicalComputationWA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WA continuous (KSG 1 Estimator)")
	case 2:
		result := continuous.MorphologicalComputationWA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func miwsContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WS Prime Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := makeW2W1A1(data, p)

	switch p.ContinuousMode {
	case 1:
		result := continuous.MorphologicalComputationWS1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WS continuous (KSG 1 Estimator)")
	case 2:
		result := continuous.MorphologicalComputationWS2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WS continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func miinContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_IN Prime Continuous Avg not implemented yet.")
}

func miwpContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Continuous Avg not implemented yet.")
}

func caContinuousAvg(p Parameters, data Data) {
	fmt.Println("CA Prime Continuous Avg not implemented yet.")
}

func uiContinuousAvg(p Parameters, data Data) {
	fmt.Println("UI Prime Continuous Avg not implemented yet.")
}

func ciContinuousAvg(p Parameters, data Data) {
	fmt.Println("CI Prime Continuous Avg not implemented yet.")
}

func misyContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Continuous Avg not implemented yet.")
}

func miaPrimeContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_A Prime Continuous Avg is not implemented yet.")
}
