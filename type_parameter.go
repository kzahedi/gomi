package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type Parameters struct {
	MeasureName       string
	Output            string
	UseContinuous     bool
	UseStateDependent bool
	Verbose           bool
	K                 int
	GlobalBins        int
	WBins             []int
	SBins             []int
	ABins             []int
	GlobalFile        string
	WIndices          []int
	SIndices          []int
	AIndices          []int
	WFile             string
	SFile             string
	AFile             string
	DFile             string
	WMin              []float64
	SMin              []float64
	AMin              []float64
	WMax              []float64
	SMax              []float64
	AMax              []float64
}

func (p Parameters) GenerateString(prefix string) string {
	s := fmt.Sprintf("%sMeasure:             %s", prefix, p.MeasureName)
	s = fmt.Sprintf("%s\n%sOutput:              %s", s, prefix, p.Output)
	s = fmt.Sprintf("%s\n%sUse state-dependent: %t", s, prefix, p.UseStateDependent)
	s = fmt.Sprintf("%s\n%sUse continuous:      %t", s, prefix, p.UseContinuous)
	s = fmt.Sprintf("%s\n%sVerbose:             %t", s, prefix, p.Verbose)
	s = fmt.Sprintf("%s\n%sk:                   %d", s, prefix, p.K)
	s = fmt.Sprintf("%s\n%sBins:                %d", s, prefix, p.GlobalBins)
	s = fmt.Sprintf("%s\n%sW bins:              %v", s, prefix, p.WBins)
	s = fmt.Sprintf("%s\n%sS bins:              %v", s, prefix, p.SBins)
	s = fmt.Sprintf("%s\n%sA bins:              %v", s, prefix, p.ABins)
	s = fmt.Sprintf("%s\n%sFull data set:       %s", s, prefix, p.GlobalFile)
	s = fmt.Sprintf("%s\n%sW indices:           %v", s, prefix, p.WIndices)
	s = fmt.Sprintf("%s\n%sS indices:           %v", s, prefix, p.SIndices)
	s = fmt.Sprintf("%s\n%sA indices:           %v", s, prefix, p.AIndices)
	s = fmt.Sprintf("%s\n%sW data set:          %s", s, prefix, p.WFile)
	s = fmt.Sprintf("%s\n%sS data set:          %s", s, prefix, p.SFile)
	s = fmt.Sprintf("%s\n%sA data set:          %s", s, prefix, p.AFile)
	s = fmt.Sprintf("%s\n%sW domains:           %v %v", s, prefix, p.WMin, p.WMax)
	s = fmt.Sprintf("%s\n%sS domains:           %v %v", s, prefix, p.SMin, p.SMax)
	s = fmt.Sprintf("%s\n%sA domains:           %v %v", s, prefix, p.AMin, p.AMax)
	return s
}

func (p Parameters) String() string {
	return p.GenerateString("")
}

func CreateParametersContainer() Parameters {
	return Parameters{MeasureName: "",
		UseContinuous:     false,
		UseStateDependent: false,
		Verbose:           false,
		K:                 0,
		GlobalBins:        0,
		WBins:             []int{},
		SBins:             []int{},
		ABins:             []int{},
		WIndices:          []int{},
		SIndices:          []int{},
		AIndices:          []int{},
		WMin:              []float64{},
		WMax:              []float64{},
		SMin:              []float64{},
		SMax:              []float64{},
		AMin:              []float64{},
		AMax:              []float64{},
		Output:            "",
		GlobalFile:        "",
		WFile:             "",
		SFile:             "",
		AFile:             ""}
}

func (p *Parameters) AddMeasureName(name string) {
	p.MeasureName = name
}

func (p *Parameters) SetUseStateDependent(b bool) {
	p.UseStateDependent = b
}

func (p *Parameters) SetUseContinuous(b bool) {
	p.UseContinuous = b
}

func (p *Parameters) AddWBins(wBins string) {
	p.WBins = parseIntString(wBins)
}

func (p *Parameters) AddSBins(wBins string) {
	p.SBins = parseIntString(wBins)
}

func (p *Parameters) AddABins(wBins string) {
	p.ABins = parseIntString(wBins)
}

func (p *Parameters) AddWIndices(wIndices string) {
	p.WIndices = parseIntString(wIndices)
}

func (p *Parameters) AddSIndices(wIndices string) {
	p.SIndices = parseIntString(wIndices)
}

func (p *Parameters) AddAIndices(wIndices string) {
	p.AIndices = parseIntString(wIndices)
}

func (p *Parameters) AddGlobalBins(bins int) {
	p.GlobalBins = bins
}

func (p *Parameters) AddGlobalFile(file string) {
	p.GlobalFile = file
}

func (p *Parameters) AddWFile(file string) {
	p.WFile = file
}

func (p *Parameters) AddSFile(file string) {
	p.SFile = file
}

func (p *Parameters) AddAFile(file string) {
	p.AFile = file
}

type T struct {
	Wmin []float64 `yaml:"W min"`
	Wmax []float64 `yaml:"W max"`
	Smin []float64 `yaml:"S min"`
	Smax []float64 `yaml:"S max"`
	Amin []float64 `yaml:"A min"`
	Amax []float64 `yaml:"A max"`
}

func (p *Parameters) AddDFile(file string) {
	p.DFile = file
	if p.DFile == "" {
		return
	}

	t := T{}

	data, err := ioutil.ReadFile(p.DFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		panic(err)
	}

	p.WMin = t.Wmin
	p.WMax = t.Wmax

	p.SMin = t.Smin
	p.SMax = t.Smax

	p.AMin = t.Amin
	p.AMax = t.Amax
}

func (p *Parameters) AddWMinMax(min []float64, max []float64) {
	p.WMin = min
	p.WMax = max
}

func (p *Parameters) AddSMinMax(min []float64, max []float64) {
	p.SMin = min
	p.SMax = max
}

func (p *Parameters) AddAMinMax(min []float64, max []float64) {
	p.AMin = min
	p.AMax = max
}

func (p *Parameters) AddK(k int) {
	p.K = k
}

func (p *Parameters) AddOutput(output string) {
	p.Output = output
}

func (p *Parameters) AddVerbose(verbose bool) {
	p.Verbose = verbose
}
