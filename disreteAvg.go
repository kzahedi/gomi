package main

import (
	"fmt"

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
	fmt.Println("MI_W Discrete Avg")

	pw2w1a1 := makePW2W1A1(data, p)
	result := discrete.MorphologicalComputationW(pw2w1a1)

	writeOutput(p, result, "MI_W")
}

func miaDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_A Discrete Avg")
	data.Discretise(p)
}

func miaPrimeDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_A Prime Discrete Avg")
	data.Discretise(p)
}

func mimiDiscreteAvg(p Parameters, data Data) {
	fmt.Println("MI_MI Prime Discrete Avg")
	data.Discretise(p)
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
