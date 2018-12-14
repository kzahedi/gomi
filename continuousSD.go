package gomi

import (
	"fmt"

	goent_c "github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/gomi/continuous/state"
)

// ContinuousSDCalculations returns the value of the selected continuous measure
// state-dependent (or point-wise)
func ContinuousSDCalculations(p Parameters, d Data) {
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
	if p.Verbose {
		fmt.Println("MI_W Continuous SD")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}
	result := state.MorphologicalComputationW(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputSD(p, result, "MI_W continuous")
}

func miaContinuousSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_A Continuous SD")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	result := state.MorphologicalComputationA(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputSD(p, result, "MI_A continuous")
}

func mimiContinuousSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_MI Prime Continuous SD")
	}

	w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices := MakeW2W1S1A1(data, p)
	if p.DFile != "" {
		w2w1s1a1 = NormaliseContinuousData(w2w1s1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.SensorMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.SensorMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1s1a1 = goent_c.Normalise(w2w1s1a1, p.Verbose)
	}

	switch p.ContinuousMode {
	case 1:
		result := state.MorphologicalComputationMI1(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_MI continuous (KSG 1 Estimator)")
	case 2:
		result := state.MorphologicalComputationMI2(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_MI continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}

}

func micaContinuousSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Continuous SD")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	switch p.ContinuousMode {
	case 1:
		result := state.MorphologicalComputationCA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_CA continuous (KSG 1 Estimator)")
	case 2:
		result := state.MorphologicalComputationCA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_CA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func miwaContinuousSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WA Continuous SD")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	switch p.ContinuousMode {
	case 1:
		result := state.MorphologicalComputationWA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_WA continuous (KSG 1 Estimator)")
	case 2:
		result := state.MorphologicalComputationWA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_WA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}

}

func miwsContinuousSD(p Parameters, data Data) {
	if p.Verbose {
		fmt.Println("MI_WS Continuous SD")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	switch p.ContinuousMode {
	case 1:
		result := state.MorphologicalComputationWS1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_WS continuous (KSG 1 Estimator)")
	case 2:
		result := state.MorphologicalComputationWS2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputSD(p, result, "MI_WS continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
}

func misyContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_SY Continuous SD")
}

func miwpContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_Wp Continuous SD")
}

func caContinuousSD(p Parameters, data Data) {
	fmt.Println("CA Continuous SD")
}

func uiContinuousSD(p Parameters, data Data) {
	fmt.Println("UI Continuous SD")
}

func ciContinuousSD(p Parameters, data Data) {
	fmt.Println("CI Continuous SD")
}

func miinContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_IN Continuous SD")
}

func miaPrimeContinuousSD(p Parameters, data Data) {
	fmt.Println("MI_A Prime Continuous SD")
}
