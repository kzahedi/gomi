package main

import (
	"fmt"

	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/utils"
)

type DataDiscretised struct {
	W [][]int
	S [][]int
	A [][]int
}

type Data struct {
	W           [][]float64
	S           [][]float64
	A           [][]float64
	Discretised DataDiscretised
}

func (d Data) String() string {
	s := ""
	if len(d.W) > 0 {
		s = fmt.Sprintf("%s\nW has %d columns and %d rows.", s, len(d.W[0]), len(d.W))
	}
	if len(d.S) > 0 {
		s = fmt.Sprintf("%s\nS has %d columns and %d rows.", s, len(d.S[0]), len(d.S))
	}
	if len(d.A) > 0 {
		s = fmt.Sprintf("%s\nA has %d columns and %d rows.", s, len(d.A[0]), len(d.A))
	}
	if len(d.Discretised.W) > 0 {
		s = fmt.Sprintf("%s\nDiscretised W has %d columns and %d rows.", s, len(d.Discretised.W[0]), len(d.Discretised.W))
	}
	if len(d.Discretised.S) > 0 {
		s = fmt.Sprintf("%s\nDiscretised S has %d columns and %d rows.", s, len(d.Discretised.S[0]), len(d.Discretised.S))
	}
	if len(d.Discretised.A) > 0 {
		s = fmt.Sprintf("%s\nDiscretised A has %d columns and %d rows.", s, len(d.Discretised.A[0]), len(d.Discretised.A))
	}
	if s == "" {
		s = "No data given."
	}
	return s
}

func (d *Data) Read(p Paramters) {
	if p.GlobalFile != "" {
		data, _ := utils.ReadFloatCsv(p.GlobalFile)
		var wdata [][]float64
		var sdata [][]float64
		var adata [][]float64
		if len(p.WIndices) > 0 {
			wdata = utils.GetFloatColumns(data, p.WIndices)
		}
		if len(p.SIndices) > 0 {
			sdata = utils.GetFloatColumns(data, p.SIndices)
		}
		if len(p.AIndices) > 0 {
			adata = utils.GetFloatColumns(data, p.AIndices)
		}
		d.W = wdata
		d.S = sdata
		d.A = adata
		return
	}

	if p.WFile != "" {
		wdata, _ := utils.ReadFloatCsv(p.WFile)
		d.W = wdata
	}

	if p.AFile != "" {
		adata, _ := utils.ReadFloatCsv(p.AFile)
		d.A = adata
	}

	if p.SFile != "" {
		sdata, _ := utils.ReadFloatCsv(p.SFile)
		d.S = sdata
	}
}

func discretiseData(data [][]float64, globalBins int) [][]int {
	min, max := dh.GetMinMax(data)
	bins := make([]int, len(min), len(min))
	for i := 0; i < len(min); i++ {
		bins[i] = globalBins
	}
	return dh.Discretise(data, bins, min, max)
}

func (d *Data) Discretise(p Paramters) {
	if p.GlobalBins > 0 {
		d.Discretised.W = discretiseData(d.W, p.GlobalBins)
		d.Discretised.S = discretiseData(d.S, p.GlobalBins)
		d.Discretised.A = discretiseData(d.A, p.GlobalBins)
		return
	}
}

func (d *Data) ClearContinuousData() {
	d.W = make([][]float64, 0, 0)
	d.S = make([][]float64, 0, 0)
	d.A = make([][]float64, 0, 0)
}
