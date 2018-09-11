package main

func calculateWBins(p Parameters, d Data) int {
	wbins := 1
	if len(p.WBins) > 0 {
		for _, v := range p.WBins {
			wbins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			wbins *= p.GlobalBins
		}
	}
	return wbins
}

func calculateABins(p Parameters, d Data) int {
	abins := 1
	if len(p.ABins) > 0 {
		for _, v := range p.ABins {
			abins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			abins *= p.GlobalBins
		}
	}
	return abins
}

func calculateSBins(p Parameters, d Data) int {
	sbins := 1
	if len(p.SBins) > 0 {
		for _, v := range p.SBins {
			sbins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			sbins *= p.GlobalBins
		}
	}
	return sbins
}
