package main

// CalculateWBins returns the number of bins for the world states depending
// on provided data
func CalculateWBins(p Parameters, d Data) int {
	wBins := 1
	if len(p.WBins) > 0 {
		for _, v := range p.WBins {
			wBins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			wBins *= p.GlobalBins
		}
	}
	return wBins
}

// CalculateABins returns the number of bins for the actuator states depending
// on provided data
func CalculateABins(p Parameters, d Data) int {
	aBins := 1
	if len(p.ABins) > 0 {
		for _, v := range p.ABins {
			aBins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			aBins *= p.GlobalBins
		}
	}
	return aBins
}

// CalculateSBins returns the number of bins for the sensor states depending
// on provided data
func CalculateSBins(p Parameters, d Data) int {
	sBins := 1
	if len(p.SBins) > 0 {
		for _, v := range p.SBins {
			sBins *= v
		}
	} else {
		for i := 0; i < len(d.W[0]); i++ {
			sBins *= p.GlobalBins
		}
	}
	return sBins
}
