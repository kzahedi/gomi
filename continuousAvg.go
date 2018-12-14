package gomi

import (
	"fmt"

	goent_c "github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/gomi/continuous"
)

// ContinuousAvgCalculations returns the averaged morphological computation
// based on estimators for continuous data. Note that not all measures are
// available on continuous state spaces.
func ContinuousAvgCalculations(p Parameters, d Data) float64 {
	var r float64
	switch p.MeasureName {
	case "MI_W":
		r = MiWContinuousAvg(p, d)
	case "MI_A":
		r = MiAContinuousAvg(p, d)
	case "MI_A_Prime":
		r = MiAPrimeContinuousAvg(p, d)
	case "MI_MI":
		r = MiMiContinuousAvg(p, d)
	case "MI_SY":
		r = MiSyContinuousAvg(p, d)
	case "MI_CA":
		r = MiCaContinuousAvg(p, d)
	case "MI_WA":
		r = MiWaContinuousAvg(p, d)
	case "MI_WS":
		r = MiWsContinuousAvg(p, d)
	case "MI_Wp":
		r = MiWpContinuousAvg(p, d)
	case "CA":
		r = CaContinuousAvg(p, d)
	case "UI":
		r = UIContinuousAvg(p, d)
	case "CI":
		r = CiContinuousAvg(p, d)
	case "MI_IN":
		r = MiInContinuousAvg(p, d)
	default:
		fmt.Println(fmt.Sprintf("unknown measure given %s in the context of continuous-avg measures.", p.MeasureName))
	}
	return r
}

// MiWContinuousAvg returns the result of the quantification MI_W. This function
// also writes the result to a file as specified in the parameters p
//    MI_W = I(W';W|A)
func MiWContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_W Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}
	result = continuous.MorphologicalComputationW(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_W continuous")
	return
}

// MiAContinuousAvg returns the result of the quantification MI_A. This function
// also writes the result to a file as specified in the parameters p
//    MI_A = I(W';A|W)
func MiAContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_A Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, a1Indices := MakeW2W1A1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	result = continuous.MorphologicalComputationA(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
	writeOutputAvg(p, result, "MI_A continuous")
	return
}

// MiAPrimeContinuousAvg returns the result of the quantification MI_A'. This function
// also writes the result to a file as specified in the parameters p
//    MI_A' = 1 - I(W';A|W)/log|W|
func MiAPrimeContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("MI_A Prime Continuous Avg is not implemented yet.")
	return -1.0
}

// MiMiContinuousAvg returns the result of the quantification MI_MI. This function
// also writes the result to a file as specified in the parameters p
//    MI_MI = I(W';W) - I(A;S)
func MiMiContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_MI Prime Continuous Avg")
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
		result = continuous.MorphologicalComputationMI1(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_MI continuous (KSG 1 Estimator)")
	case 2:
		result = continuous.MorphologicalComputationMI2(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_MI continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
	return
}

// MiCaContinuousAvg returns the result of the quantification MI_CA. This function
// also writes the result to a file as specified in the parameters p
//    MI_CA = I(W';W) - I(W';A)
func MiCaContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_CA Continuous Avg")
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
		result = continuous.MorphologicalComputationCA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_CA continuous (KSG 1 Estimator)")
	case 2:
		result = continuous.MorphologicalComputationCA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_CA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
	return
}

// MiWaContinuousAvg returns the result of the quantification MI_WA. This function
// also writes the result to a file as specified in the parameters p
//    MI_WA = I(W;{W,A}) - I(W';A)
func MiWaContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_WA Continuous Avg")
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
		result = continuous.MorphologicalComputationWA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WA continuous (KSG 1 Estimator)")
	case 2:
		result = continuous.MorphologicalComputationWA2(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WA continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
	return
}

// MiWsContinuousAvg returns the result of the quantification MI_WS. This function
// also writes the result to a file as specified in the parameters p
//    MI_WS = I(W;{W,S}) - I(W';S)
func MiWsContinuousAvg(p Parameters, data Data) (result float64) {
	if p.Verbose {
		fmt.Println("MI_WS Prime Continuous Avg")
	}

	w2w1a1, w2Indices, w1Indices, s1Indices := MakeW2W1S1(data, p)
	if p.DFile != "" {
		w2w1a1 = NormaliseContinuousData(w2w1a1,
			[][]float64{p.WorldMin, p.WorldMin, p.ActuatorMin},
			[][]float64{p.WorldMax, p.WorldMax, p.ActuatorMax}, p.Verbose)
	} else {
		w2w1a1 = goent_c.Normalise(w2w1a1, p.Verbose)
	}

	switch p.ContinuousMode {
	case 1:
		result = continuous.MorphologicalComputationWS1(w2w1a1, w2Indices, w1Indices, s1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WS continuous (KSG 1 Estimator)")
	case 2:
		result = continuous.MorphologicalComputationWS2(w2w1a1, w2Indices, w1Indices, s1Indices, p.K, p.Verbose)
		writeOutputAvg(p, result, "MI_WS continuous (KSG 2 Estimator)")
	default:
		fmt.Println(fmt.Sprintf("Unknown Continuous Mode %d", p.ContinuousMode))
	}
	return
}

// MiInContinuousAvg returns the result of the quantification MI_IN. This function
// also writes the result to a file as specified in the parameters p
//    MI_IN = log|A| - I(A;S)
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func MiInContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("MI_IN Prime Continuous Avg not implemented yet.")
	return -1.0
}

// MiWpContinuousAvg returns the result of the quantification MI_Wp. This function
// also writes the result to a file as specified in the parameters p
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func MiWpContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("MI_Wp Prime Continuous Avg not implemented yet.")
	return -1.0
}

// CaContinuousAvg returns the result of the quantification CA. This function
// also writes the result to a file as specified in the parameters p
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func CaContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("CA Prime Continuous Avg not implemented yet.")
	return -1.0
}

// UIContinuousAvg returns the result of the quantification UI. This function
// also writes the result to a file as specified in the parameters p
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func UIContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("UI Prime Continuous Avg not implemented yet.")
	return -1.0
}

// CiContinuousAvg returns the result of the quantification CI. This function
// also writes the result to a file as specified in the parameters p
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func CiContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("CI Prime Continuous Avg not implemented yet.")
	return -1.0
}

// MiSyContinuousAvg returns the result of the quantification MI_SY. This function
// also writes the result to a file as specified in the parameters p
// This function returns -1.0, because the quantification is not implemented
// based on entropy estimators for continuous state spaces yet.
func MiSyContinuousAvg(p Parameters, data Data) float64 {
	fmt.Println("MI_SY Prime Continuous Avg not implemented yet.")
	return -1.0
}
