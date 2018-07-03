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

	w2w1a1, w2indices, w1indices, a1indices := makeW2W1A1(data, p)
	result := continuous.MorphologicalComputationW(w2w1a1, w2indices, w1indices, a1indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_W continuous")

}

func miaContinuousAvg(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Continuous Avg")
	}

	w2w1a1, w2indices, w1indices, a1indices := makeW2W1A1(data, p)
	result := continuous.MorphologicalComputationA(w2w1a1, w2indices, w1indices, a1indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_W continuous")
}

func miaPrimeContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_A Prime Continuous Avg is not implemented yet.")
}

func mimiContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_MI Prime Continuous Avg")
}

func misyContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Continuous Avg")
}

func micaContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_CA Prime Continuous Avg")
}

func miwaContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_WA Prime Continuous Avg")
}

func miwsContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_WS Prime Continuous Avg")
}

func miwpContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Continuous Avg")
}

func caContinuousAvg(p Parameters, data Data) {
	fmt.Println("CA Prime Continuous Avg")
}

func uiContinuousAvg(p Parameters, data Data) {
	fmt.Println("UI Prime Continuous Avg")
}

func ciContinuousAvg(p Parameters, data Data) {
	fmt.Println("CI Prime Continuous Avg")
}

func miinContinuousAvg(p Parameters, data Data) {
	fmt.Println("MI_IN Prime Continuous Avg")
}
