package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kzahedi/goent/dh"
	entropy "github.com/kzahedi/goent/discrete"
)

func makePW2W1A1(d Data, p Parameters) [][][]float64 {
	var wbins []int
	var abins []int

	d.Discretise(p)

	if len(p.WBins) > 0 {
		wbins = p.WBins
	} else {
		for i := 0; i < len(d.Discretised.W[0]); i++ {
			wbins = append(wbins, p.GlobalBins)
		}
	}

	if len(p.ABins) > 0 {
		wbins = p.ABins
	} else {
		for i := 0; i < len(d.Discretised.A[0]); i++ {
			abins = append(abins, p.GlobalBins)
		}
	}

	w := dh.MakeUnivariateRelabelled(d.Discretised.W, wbins)
	a := dh.MakeUnivariateRelabelled(d.Discretised.A, abins)

	w2w1a1 := make([][]int, len(w)-1, len(w)-1)

	for i := 0; i < len(w)-1; i++ {
		w2w1a1[i] = make([]int, 3, 3)
		w2w1a1[i][0] = w[i+1]
		w2w1a1[i][1] = w[i]
		w2w1a1[i][2] = a[i]
	}

	pw2w1a1 := entropy.Emperical3D(w2w1a1)
	return pw2w1a1
}

func writeOutput(p Parameters, result float64, label string) {
	str := fmt.Sprintf("%s: %f", label, result)

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
