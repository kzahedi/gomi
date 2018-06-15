package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	entropy "github.com/kzahedi/goent/discrete"
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
	data.Discretise(p)

	w2w1a1 := makeW2W1A1(data, p)
	pw2w1a1 := entropy.Emperical3D(w2w1a1)
	miw := discrete.MorphologicalComputationW(pw2w1a1)

	str := fmt.Sprintf("MI_W: %f", miw)

	if p.Verbose {
		fmt.Println(str)
	}

	file, err := os.Create(p.Output)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(str)

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
