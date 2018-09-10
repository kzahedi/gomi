package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v1"
)

type Parameters struct {
	MeasureName       string
	Output            string
	ConfigFile        string
	UseContinuous     bool
	UseStateDependent bool
	Verbose           bool
	ContinuousMode    int
	K                 int
	GlobalBins        int
	Iterations        int
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
	s = fmt.Sprintf("%s\n%sConfig file:         %s", s, prefix, p.ConfigFile)
	s = fmt.Sprintf("%s\n%sUse state-dependent: %t", s, prefix, p.UseStateDependent)
	s = fmt.Sprintf("%s\n%sUse continuous:      %t", s, prefix, p.UseContinuous)
	s = fmt.Sprintf("%s\n%sVerbose:             %t", s, prefix, p.Verbose)
	s = fmt.Sprintf("%s\n%sContinuous Mode:     %d", s, prefix, p.ContinuousMode)
	s = fmt.Sprintf("%s\n%sk:                   %d", s, prefix, p.K)
	s = fmt.Sprintf("%s\n%sBins:                %d", s, prefix, p.GlobalBins)
	s = fmt.Sprintf("%s\n%sIterations:          %d", s, prefix, p.Iterations)
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
		Iterations:        0,
		ContinuousMode:    0,
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
		ConfigFile:        "",
		GlobalFile:        "",
		WFile:             "",
		SFile:             "",
		AFile:             ""}
}

func (p *Parameters) SetMeasureName(name string) {
	p.MeasureName = name
}

func (p *Parameters) SetUseStateDependent(b bool) {
	p.UseStateDependent = b
}

func (p *Parameters) SetUseContinuous(b bool) {
	p.UseContinuous = b
}

func (p *Parameters) SetWBins(wBins string) {
	p.WBins = parseIntString(wBins)
}

func (p *Parameters) SetSBins(wBins string) {
	p.SBins = parseIntString(wBins)
}

func (p *Parameters) SetABins(wBins string) {
	p.ABins = parseIntString(wBins)
}

func (p *Parameters) SetWIndices(wIndices string) {
	p.WIndices = parseIntString(wIndices)
}

func (p *Parameters) SetSIndices(wIndices string) {
	p.SIndices = parseIntString(wIndices)
}

func (p *Parameters) SetAIndices(wIndices string) {
	p.AIndices = parseIntString(wIndices)
}

func (p *Parameters) SetGlobalBins(bins int) {
	p.GlobalBins = bins
}

func (p *Parameters) SetGlobalFile(file string) {
	p.GlobalFile = file
}

func (p *Parameters) SetWFile(file string) {
	p.WFile = file
}

func (p *Parameters) SetSFile(file string) {
	p.SFile = file
}

func (p *Parameters) SetAFile(file string) {
	p.AFile = file
}

func (p *Parameters) SetIterations(iterations int) {
	p.Iterations = iterations
}

type T struct {
	Wmin []float64 `yaml:"W min"`
	Wmax []float64 `yaml:"W max"`
	Smin []float64 `yaml:"S min"`
	Smax []float64 `yaml:"S max"`
	Amin []float64 `yaml:"A min"`
	Amax []float64 `yaml:"A max"`
}

func (p *Parameters) SetDFile(file string) {
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

func (p *Parameters) SetWMinMax(min []float64, max []float64) {
	p.WMin = min
	p.WMax = max
}

func (p *Parameters) SetSMinMax(min []float64, max []float64) {
	p.SMin = min
	p.SMax = max
}

func (p *Parameters) SetAMinMax(min []float64, max []float64) {
	p.AMin = min
	p.AMax = max
}

func (p *Parameters) SetK(k int) {
	p.K = k
}

func (p *Parameters) SetOutput(output string) {
	p.Output = output
}

func (p *Parameters) SetVerbose(verbose bool) {
	p.Verbose = verbose
}

func (p *Parameters) SetContinuousMode(cm int) {
	p.ContinuousMode = cm
}

type CfgT struct {
	Measure        string `yaml:"Measure"`
	Continuous     bool   `yaml:"Continuous"`
	ContinuousMode int    `yaml:"Continuous mode"`
	UseState       bool   `yaml:"State-dependent"`
	Verbose        bool   `yaml:"Verbose"`
	Bins           int    `yaml:"Bins"`
	Iterations     int    `yaml:"Iterations"`
	K              int    `yaml:"k"`
	Output         string `yaml:"Output file"`
	WBins          string `yaml:"W Bins"`
	ABins          string `yaml:"A Bins"`
	SBins          string `yaml:"S Bins"`
	WIndices       string `yaml:"W Indices"`
	AIndices       string `yaml:"A Indices"`
	SIndices       string `yaml:"S Indices"`
	File           string `yaml:"Full data file"`
	WFile          string `yaml:"W data file"`
	AFile          string `yaml:"A data file"`
	SFile          string `yaml:"S data file"`
	DFile          string `yaml:"Domain file"`
}

func (p *Parameters) SetConfigFile(file string) {
	p.ConfigFile = file
	if p.ConfigFile == "" {
		return
	}

	t := CfgT{}

	data, err := ioutil.ReadFile(p.ConfigFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		panic(err)
	}

	p.SetMeasureName(t.Measure)
	p.SetUseContinuous(t.Continuous)
	p.SetContinuousMode(t.ContinuousMode)
	p.SetUseStateDependent(t.UseState)
	p.SetGlobalBins(t.Bins)
	p.SetWBins(t.WBins)
	p.SetSBins(t.SBins)
	p.SetABins(t.ABins)
	p.SetK(t.K)
	p.SetOutput(t.Output)
	p.SetVerbose(t.Verbose)
	p.SetGlobalFile(t.File)
	p.SetWIndices(t.WIndices)
	p.SetSIndices(t.SIndices)
	p.SetAIndices(t.AIndices)
	p.SetWFile(t.WFile)
	p.SetSFile(t.SFile)
	p.SetAFile(t.AFile)
	p.SetDFile(t.DFile)
	p.SetIterations(t.Iterations)

	p.Verbose = true
}

func (p *Parameters) CheckParameters() {
	errorMsg := ""
	hasError := false
	if _, err := os.Stat(p.File); os.IsNotExist(err) {
		hasError = true
		errorMsg = fmt.Sprintf("File %f not found\n", p.File)
	}
	if _, err := os.Stat(p.WFile); os.IsNotExist(err) {
		hasError = true
		errorMsg = fmt.Sprintf("%sW File %f not found\n", p.WFile)
	}
	if _, err := os.Stat(p.SFile); os.IsNotExist(err) {
		hasError = true
		errorMsg = fmt.Sprintf("%sS File %f not found\n", p.SFile)
	}
	if _, err := os.Stat(p.AFile); os.IsNotExist(err) {
		hasError = true
		errorMsg = fmt.Sprintf("%sA File %f not found\n", p.AFile)
	}

	if hasError == true {
		fmt.Println(errorMsg)
		os.Exit(-1)
	}
}
