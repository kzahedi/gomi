package main

import "fmt"

func continuousSDCalculations(p Parameters, d Data) {
	switch p.MeasureName {
	case "MI_W":
		miwContinuousSD(p, d)
	case "MI_A":
		miaContinuousSD(p, d)
	case "MI_A_Prime":
		miaPrimeContinuousSD(p, d)
	case "MI_MI":
		mimiContinuousSD(p, d)
	case "MI_SY":
		misyContinuousSD(p, d)
	case "MI_SY_NID":
		misynidContinuousSD(p, d)
	case "MI_CA":
		micaContinuousSD(p, d)
	case "MI_WA":
		miwaContinuousSD(p, d)
	case "MI_WS":
		miwsContinuousSD(p, d)
	case "MI_Wp":
		miwpContinuousSD(p, d)
	case "CA":
		caContinuousSD(p, d)
	case "UI":
		uiContinuousSD(p, d)
	case "CI":
		ciContinuousSD(p, d)
	case "MI_IN":
		miinContinuousSD(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of continuous-state-dependent measures.", p.MeasureName))
	}
}

func miwContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_W Continuous SD")
	data.Discretise(p)
}

func miaContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_A Continuous SD")
	data.Discretise(p)
}

func miaPrimeContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_A Prime Continuous SD")
	data.Discretise(p)
}

func mimiContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_MI Prime Continuous SD")
	data.Discretise(p)
}

func misyContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_SY Prime Continuous SD")
	data.Discretise(p)
}

func misynidContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_SY_NID Prime Continuous SD")
	data.Discretise(p)
}

func micaContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_CA Prime Continuous SD")
	data.Discretise(p)
}

func miwaContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_WA Prime Continuous SD")
	data.Discretise(p)
}

func miwsContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_WS Prime Continuous SD")
	data.Discretise(p)
}

func miwpContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_Wp Prime Continuous SD")
	data.Discretise(p)
}

func caContinuousSD(p Parameters, data Data) {
	fmt.Println("CA Prime Continuous SD")
	data.Discretise(p)
}

func uiContinuousSD(p Parameters, data Data) {
	fmt.Println("UI Prime Continuous SD")
	data.Discretise(p)
}

func ciContinuousSD(p Parameters, data Data) {
	fmt.Println("CI Prime Continuous SD")
	data.Discretise(p)
}

func miinContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_IN Prime Continuous SD")
	data.Discretise(p)
}
